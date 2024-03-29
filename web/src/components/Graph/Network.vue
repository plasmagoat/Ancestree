<template>
  <div id="network" class="relative">
    <div
      v-show="linkTextVisible"
      class="linkText"
      :style="linkTextPosition"
      v-text="linkTextContent"
    />
    <svg
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      :width="svgSize.width"
      :height="svgSize.height"
      :viewBox="viewBox"
      class="bg-transparent"
      overflow="visible"
      @click="clickEle"
      @mouseover.prevent="svgMouseover"
      @mouseout="svgMouseout"
    >
      <g id="container">
        <!-- links and link-text 注：先绘制边 -->
        <g>
          <defs>
            <marker
              id="arrowhead"
              markerWidth="15"
              markerHeight="7"
              refX="0"
              refY="3.5"
              orient="auto"
            >
              <polygon points="0 0, 15 3.5, 0 7" />
            </marker>
            <marker
              id="arrow-head"
              viewBox="0 -5 10 10"
              refX="30"
              refY="-1"
              markerWidth="4"
              markerHeight="4"
              orient="auto"
              transform="rotate(-20 0 0)"
            >
              <path d="M0,-5L10,0L0,5" />
            </marker>
          </defs>
          <g
            :key="link.index"
            v-for="link in links"
            fill="none"
            :stroke-width="linkWidth"
          >
            <line
              v-if="!linkArc"
              :class="
                `${link[linkTypeKey]} ${link.selected} link element stroke-current fill`
              "
              :stroke-width="linkWidth"
              marker-start="url(#arrowhead)"
            />
            <path
              v-if="linkArc"
              :class="
                `${link[linkTypeKey]} ${link.selected} linkarc element stroke-current fill`
              "
              :d="
                `M${link.source.x},${link.source.y} A${Math.hypot(
                  link.target.x - link.source.x,
                  link.target.y - link.source.y,
                )},${Math.hypot(
                  link.target.x - link.source.x,
                  link.target.y - link.source.y,
                )} 0 0,1 ${link.target.x},${link.target.y}`
              "
              marker-end="url(#arrow-head)"
            />
            <!-- dx dy 文字左下角坐标 -->
            <text
              v-if="showLinkText"
              dx="0"
              dy="0"
              class="link-text"
              :fill="theme.textFill"
              :font-size="linkTextFrontSize"
            >
              {{ link[linkTextKey] }}
            </text>
          </g>
        </g>

        <!-- node and node-text -->
        <g id="node-group">
          <g :key="node.index" v-for="node in nodes">
            <circle
              :fill="nodeColor(node[nodeTypeKey])"
              :stroke-width="highlightNodes.indexOf(node.id) == -1 ? 3 : 4"
              :stroke="
                highlightNodes.indexOf(node.id) == -1
                  ? theme.nodeStroke
                  : 'gold'
              "
              :class="
                `${node[nodeTypeKey]} ${
                  node.showText ? 'selected' : ''
                } node element`
              "
              :r="nodeSize"
            />
            <text
              v-show="node.showText"
              :dx="nodeSize + 5"
              dy="0"
              class="node-text fill-current dark:text-gray-200"
              :fill="theme.textFill"
              :font-size="nodeTextFontSize"
            >
              {{ node[nodeTextKey] }}
            </text>
          </g>
          <g />
        </g>
      </g>
    </svg>
    <div class="absolute w-full bottom-1">
      <input
        class="relative block mx-auto rounded-lg overflow-hidden appearance-none bg-gray-400 h-3 w-128"
        type="range"
        min="0.1"
        max="4"
        step="0.1"
        :value="scale"
        @change="slider"
        @input="slider"
      />
    </div>
  </div>
</template>

<script>
import * as d3 from 'd3'
// import * as d3Force from 'd3-force'
// import * as d3Zoom from 'd3-zoom'
// import * as d3Scale from 'd3-scale'
// import * as d3Selection from 'd3-selection'
// import * as d3Drag from 'd3-drag'
// import * as d3ScaleChromatic from 'd3-scale-chromatic'
// import d3SelectionMulti from "d3-selection-multi";
// const d3 = Object.assign({}, d3Force, d3Zoom, d3Scale, d3Selection, d3Drag)
// 元素的 classList 是 DOMTokenList
DOMTokenList.prototype.indexOf = Array.prototype.indexOf
export default {
  name: 'network',
  props: {
    nodeList: Array,
    linkList: Array,
    // node
    nodeSize: {
      type: Number,
      default: 14,
    },
    nodeTextKey: {
      type: String,
      default: 'id',
    },
    nodeTypeKey: {
      type: String,
      default: 'gender',
    },
    nodeTextFontSize: {
      type: Number,
      default: 14,
    },
    // link
    linkWidth: {
      type: Number,
      default: 2,
    },
    linkArc: {
      type: Boolean,
      default: false,
    },
    showLinkText: {
      type: Boolean,
      default: false,
    },
    linkTextKey: {
      type: String,
      default: 'value',
    },
    linkTypeKey: {
      type: String,
      default: 'type',
    },
    linkTextFrontSize: {
      type: Number,
      default: 10,
    },
    linkStrength: {
      type: Number,
      default: 0.2,
    },
    linkDistance: {
      type: Number,
      default: 50,
    },
    // svg
    svgSize: {
      type: Object,
      default: () => {
        return {
          width: window.innerWidth,
          height: window.innerHeight,
        }
      },
    },
    svgTheme: {
      type: String,
      default: 'light', // dark or light
    },
    bodyStrength: {
      type: Number,
      default: -200,
    },
    // others
    highlightNodes: {
      type: Array,
      default: () => {
        return []
      },
    },
  },
  data() {
    return {
      selection: {
        links: [],
        nodes: [],
      },
      pinned: [], // 被订住的节点的下标
      force: null,
      zoom: d3.zoom(),
      scale: 1,
      nodeColor: d3.scaleOrdinal(d3.schemeCategory10),
      linkTextVisible: false,
      linkTextPosition: {
        top: 0,
        left: 0,
      },
      linkTextContent: '',
      viewBox: [
        -window.innerWidth / 4,
        -window.innerHeight / 4,
        window.innerWidth / 2,
        window.innerHeight / 2,
      ],
    }
  },
  computed: {
    nodes() {
      // 去重
      let nodes = this.nodeList.slice()
      let nodeIds = []
      nodes = nodes.filter(node => {
        if (nodeIds.indexOf(node.id) === -1) {
          nodeIds.push(node.id)
          return true
        } else {
          return false
        }
      })
      return nodes
    },
    links() {
      return this.linkList
    },
    theme() {
      if (this.svgTheme === 'light') {
        return {
          nodeStroke: 'white',
          textFill: 'black',
        }
      } else {
        // dark
        return {
          nodeStroke: 'white',
          textFill: 'white',
        }
      }
    },
  },
  watch: {
    svgSize: function() {
      this.initData()
      this.$nextTick(function() {
        this.initDragTickZoom()
      })
    },
    bodyStrength: function() {
      this.initData()
      this.$nextTick(function() {
        this.initDragTickZoom()
      })
    },
    linkDistance: function() {
      this.initData()
      this.$nextTick(function() {
        this.initDragTickZoom()
      })
    },
    linkStrength: function() {
      this.initData()
      this.$nextTick(function() {
        this.initDragTickZoom()
      })
    },
    nodes: function() {
      this.initData()
      this.$nextTick(function() {
        // 以下这个函数重新在 node 上调用了拖拽
        // 只有在 mounted 后才有用
        // 所以要使用 $nextTick
        this.initDragTickZoom()
      })
    },
  },
  created() {
    this.initData()
  },
  mounted() {
    this.initDragTickZoom()
  },
  methods: {
    initData() {
      this.force = d3
        .forceSimulation(this.nodes)
        .alphaDecay(0.001)
        .force(
          'link',
          d3
            .forceLink(this.links)
            .id(d => d.id)
            .distance(this.linkDistance)
            .strength(this.linkStrength),
        )
        .force('charge', d3.forceManyBody().strength(this.bodyStrength)) //The strength of the attraction or repulsion
        .force('x', d3.forceX())
        .force('y', d3.forceY())
    },
    initDragTickZoom() {
      d3.selectAll('.node').call(this.drag(this.force))
      this.force.on('tick', () => {
        d3.selectAll('.link')
          .data(this.links)
          .attr('x1', d => d.source.x)
          .attr('y1', d => d.source.y)
          .attr('x2', d => d.target.x)
          .attr('y2', d => d.target.y)
        d3.selectAll('.linkarc')
          .data(this.links)
          .attr('d', this.calcLinkArc)
        d3.selectAll('.node')
          .data(this.nodes)
          .attr('cx', d => d.x)
          .attr('cy', d => d.y)
        d3.selectAll('.node-text')
          .data(this.nodes)
          .attr('x', d => d.x)
          .attr('y', d => d.y)
        d3.selectAll('.link-text')
          .data(this.links)
          .attr('x', d => (d.source.x + d.target.x) / 2)
          .attr('y', d => (d.source.y + d.target.y) / 2)
      })
      this.zoom.scaleExtent([0.1, 4]).on('zoom', this.zoomed)
      d3.select('svg')
        .call(this.zoom)
        .on('dblclick.zoom', null)
    },
    clickLink(e) {
      this.$emit('clickLink', e, e.target.__data__)
    },
    clickNode(e) {
      if (this.pinned.length === 0) {
        //this.pinnedState(e)
      } else {
        d3.selectAll('.element').style('opacity', 0.2)
        this.pinned = []
      }
      this.$emit('clickNode', e, e.target.__data__)
    },
    clickEle(e) {
      if (e.target.tagName === 'circle') {
        this.clickNode(e)
      } else if (e.target.tagName === 'line') {
        this.clickLink(e)
      }
    },
    svgMouseover(e) {
      if (e.target.nodeName === 'circle') {
        if (this.pinned.length === 0) {
          this.selectedState(e)
        }
        this.$forceUpdate()
        this.$emit('hoverNode', e, e.target.__data__)
      } else if (e.target.nodeName === 'line') {
        this.linkTextPosition = {
          left: e.clientX + 'px',
          top: e.clientY - 50 + 'px',
        }
        this.linkTextContent = e.target.__data__[this.linkTextKey]
        this.linkTextVisible = true
        this.$emit('hoverLink', e, e.target.__data__)
      }
    },
    svgMouseout(e) {
      this.linkTextVisible = false
      if (e.target.nodeName === 'circle') {
        if (this.pinned.length === 0) {
          this.noSelectedState(e)
        }
        this.$forceUpdate()
      }
    },
    selectedState(e) {
      e.target.__data__.showText = true
      e.target.classList.add('selected')
      this.selection.nodes.push(e.target.__data__)
      // 周围节点显示文字、边和结点增加 selected class、添加进 selection
      this.lightNeibor(e.target.__data__)
      // 除了 selected 的其余节点透明度减小
      //d3.selectAll(".element").style("opacity", 0.2);
    },
    noSelectedState(e) {
      // 节点自身不显示文字、移除 selected class
      e.target.__data__.showText = false
      // e.target.classList.remove("selected");
      // 周围节点不显示文字、边和结点移除 selected class
      this.darkenNerbor()
      // 所有节点透明度恢复
      d3.selectAll('.element').style('opacity', 1)
    },
    pinnedState(e) {
      this.pinned.push(e.target.__data__.index)
      d3.selectAll('.element').style('opacity', 0.05)
    },
    lightNeibor(node) {
      this.links.forEach(link => {
        if (link.source.index === node.index) {
          link.selected = 'selected'
          this.selection.links.push(link)
          this.selection.nodes.push(link.target)
          this.lightNode(link.target)
        }
        if (link.target.index === node.index) {
          link.selected = 'selected'
          this.selection.links.push(link)
          this.selection.nodes.push(link.source)
          this.lightNode(link.source)
        }
      })
    },
    lightNode(selectedNode) {
      this.nodes.forEach(node => {
        if (node.index === selectedNode.index) {
          node.showText = true
        }
      })
    },
    darkenNerbor() {
      this.links.forEach(link => {
        this.selection.links.forEach(selectedLink => {
          if (selectedLink.id === link.id) {
            link.selected = ''
          }
        })
      })
      this.nodes.forEach(node => {
        this.selection.nodes.forEach(selectednode => {
          if (selectednode.id === node.id) {
            node.showText = false
          }
        })
      })
      // 清空 selection
      this.selection.nodes = []
      this.selection.links = []
    },
    zoomed(event) {
      this.setScale(event.transform.k)
    },
    slider(event) {
      this.setScale(event.target.value)
    },
    setScale(scale) {
      this.scale = scale
      d3.select('#container').attr(
        'transform',
        `scale(${this.scale})`,
        //translate(${event.transform.x},${event.transform.y})
      )
    },
    drag(simulation) {
      function dragstarted(event, d) {
        if (!event.active) simulation.alphaTarget(0.3).restart()
        d.fx = d.x
        d.fy = d.y
      }
      function dragged(event, d) {
        d.fx = event.x
        d.fy = event.y
      }
      function dragended(event, d) {
        if (!event.active) simulation.alphaTarget(0)
        d.fx = null
        d.fy = null
      }
      return d3
        .drag()
        .on('start', dragstarted)
        .on('drag', dragged)
        .on('end', dragended)
    },
    calcLinkArc(d) {
      const r = Math.hypot(d.target.x - d.source.x, d.target.y - d.source.y)
      return `M${d.source.x},${d.source.y} A${r},${r} 0 0,1 ${d.target.x},${d.target.y}`
    },
  },
}
</script>
<style scoped>
svg {
  /* border-radius: 5px; */
}
text {
  pointer-events: none;
  /* @apply dark:text-gray-100; */
}
.element {
  transition: opacity 0.5s ease;
}
.selected {
  opacity: 1 !important;
}
.node,
.link {
  cursor: pointer;
}
.linkText {
  position: absolute;
  z-index: 10;
  background-color: rgba(75, 75, 75, 0.596);
  border-radius: 10px;
  color: white;
  padding: 10px;
}
</style>
