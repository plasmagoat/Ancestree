<template>
  <div class="py-6 flex flex-col justify-start">
    <div class="relative p-3">
      <div
        class="absolute m-2 inset-0 bg-gradient-to-r from-green-800 to-green-600 shadow-lg transform skew-y-0 -rotate-6 rounded-3xl"
      />
      <div
        class="relative px-4 py-10 bg-gray-100 dark:bg-gray-700 shadow-lg rounded-3xl border-green-800 border-2"
      >
        <h1>
          {{ node.fullname }}
        </h1>
        <h2>
          {{
            new Date(Date.parse(node.birthday)).toLocaleDateString('en-ZA', {
              weekday: 'long',
              year: 'numeric',
              month: 'long',
              day: 'numeric',
            })
          }}
        </h2>
      </div>
    </div>
    <div v-show="!editorOpen" class="p-3 flex flex-row mt-5">
      <standard-button
        class="flex-1"
        text="Add Parent"
        editmode="parent"
        @click="openEditor"
      />
      <standard-button
        class="flex-1 mx-2"
        text="Add Child"
        editmode="child"
        @click="openEditor"
      />
      <standard-button
        class="flex-1"
        text="Edit"
        editmode="edit"
        @click="openEditor"
      />
    </div>
    <div v-show="editorOpen" class="relative p-3 mt-5">
      <div
        class="absolute m-2 inset-0 bg-gradient-to-r from-green-800 to-green-600 shadow-lg transform skew-y-0 -rotate-3 rounded-3xl"
      />
      <div
        class="relative p-4 bg-gray-100 dark:bg-gray-700 shadow-lg rounded-3xl border-green-800 border-2"
      >
        <div v-show="editMode != 'edit'" class="grid grid-cols-1 gap-6">
          <person-selector
            text="Link Person"
            v-model:selected="linkingPerson"
          />
          <div class="block">
            <standard-button
              submitmode="submit"
              class="flex-1"
              text="Link"
              @click="closeLinker"
            />
          </div>
        </div>

        <div class="grid grid-cols-1 gap-6">
          <hr class="dashed" />
          <label class="block">
            <span class="text-gray-700">Full name</span>
            <input
              type="text"
              class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 dark:bg-gray-700"
              placeholder="Name"
              v-model="editingPerson.fullname"
            />
          </label>
          <date-input
            v-show="editingPerson != null"
            v-model:date="editingPerson.birthday"
            text="Birthday"
          />
          <gender-selector
            v-show="editingPerson != null"
            text="Gender"
            v-model:gender="editingPerson.gender"
          />
          <div class="block">
            <standard-button
              submitmode="submit"
              class="flex-1"
              text="Submit"
              @click="closeEditor"
            />
            <standard-button
              submitmode="cancel"
              class="flex-1 mx-2"
              text="Cancel"
              @click="resetEditor"
            />
            <standard-button
              v-show="editMode == 'edit'"
              submitmode="cancel"
              class="flex-1"
              text="Delete"
              @click="deleteSelected"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { ref, watch } from 'vue'
import StandardButton from '../Shared/FormElements/Buttons/StandardButton.vue'
import personApi from '../../api/person'
import { useStore } from 'vuex'
import GenderSelector from '../Shared/FormElements/GenderSelector.vue'
import DateInput from '../Shared/FormElements/DateInput.vue'
import PersonSelector from '../Shared/FormElements/PersonSelector.vue'
export default {
  name: 'profile',
  components: { StandardButton, GenderSelector, DateInput, PersonSelector },
  props: {
    node: Object,
  },
  setup(props) {
    const store = useStore()
    const editMode = ref(null)
    const isModalVisible = ref(false)

    const closeModal = () => {
      isModalVisible.value = false
    }
    const openModal = () => {
      isModalVisible.value = true
    }
    watch(
      () => props.node,
      () => {
        resetEditor()
      },
    )

    const editorOpen = ref(false)
    const editingPerson = ref({})
    const linkingPerson = ref({})
    const openEditor = e => {
      resetEditor()
      editMode.value = e.target.getAttribute('editmode')
      if (editMode.value == 'edit') {
        editingPerson.value = {
          ...props.node,
        }
      }
      editorOpen.value = true
    }
    const closeEditor = () => {
      switch (editMode.value) {
        case 'parent':
          personApi.createParent(props.node.id, editingPerson.value, () => {
            // update full tree for now..
            store.dispatch('graph/refreshGraph')
          })
          break
        case 'child':
          personApi.createChild(props.node.id, editingPerson.value, () => {
            // update full tree for now..
            store.dispatch('graph/refreshGraph')
          })
          break
        case 'edit':
          personApi.updatePerson(props.node.id, editingPerson.value, () => {
            // update full tree for now..
            store.dispatch('graph/refreshGraph')
          })
          break

        default:
          break
      }
      resetEditor()
    }

    const deleteSelected = () => {
      personApi.deletePerson(props.node.id, () => {})
    }

    const closeLinker = () => {
      switch (editMode.value) {
        case 'parent':
          personApi.createLink(props.node.id, linkingPerson.value, () => {
            // update full tree for now..
            store.dispatch('graph/refreshGraph')
          })
          break
        case 'child':
          personApi.createLink(linkingPerson.value, props.node.id, () => {
            // update full tree for now..
            store.dispatch('graph/refreshGraph')
          })
          break
        default:
          break
      }
      resetEditor()
    }

    const resetEditor = () => {
      editorOpen.value = false
      editingPerson.value = {}
      editMode.value = null
    }

    return {
      isModalVisible,
      closeModal,
      openModal,

      editMode,
      editorOpen,
      editingPerson,
      linkingPerson,
      openEditor,
      closeEditor,
      closeLinker,
      deleteSelected,
      resetEditor,
    }
  },
}
</script>
<style>
.text {
  @apply dark:text-gray-100;
}
</style>
