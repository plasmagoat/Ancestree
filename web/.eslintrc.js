module.exports = {
    root: true,
    env: {
        node: true
    },
    extends: ["plugin:vue/vue3-essential", "eslint:recommended", "@vue/prettier"],
    parserOptions: {
        parser: "babel-eslint"
    },
    rules: {
        'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'prettier/prettier': [
            'error',
            {
                semi: false,
                singleQuote: true,
                trailingComma: 'all',
                htmlWhitespaceSensitivity: 'ignore',
            },
        ],
        'vue/attribute-hyphenation': ['error', 'always'],
        'vue/attributes-order': [
            2,
            { order: [
                'GLOBAL', // id
                'DEFINITION', // is
                'UNIQUE', // ref, key, slot
                'CONDITIONALS', // v-if, v-else-if, v-else, v-show, v-cloak
                'LIST_RENDERING', // v-for
                'OTHER_ATTR', // all unspecified bound & unbound attributes
                // 'BINDING', // v-model
                'RENDER_MODIFIERS', // v-pre, v-once
                'CONTENT', // v-html, v-text
                'EVENTS', // v-on
            ]},
        ],
        'vue/component-name-in-template-casing': ['error', 'kebab-case'],
        'vue/html-end-tags': 'error',
        'vue/html-self-closing': [
            'error',
            {
                html: {
                    void: 'always',
                    normal: 'always',
                    component: 'always',
                },
                svg: 'always',
                math: 'always',
            },
        ],
        'vue/name-property-casing': ['error', 'kebab-case'],
        'vue/no-multi-spaces': 'error',
        'vue/order-in-components': [
            'error',
            {
                order: [
                    'el',
                    'name',
                    'parent',
                    'functional',
                    ['delimiters', 'comments'],
                    ['components', 'directives', 'filters'],
                    'extends',
                    'mixins',
                    'inheritAttrs',
                    'model',
                    ['props', 'propsData'],
                    'data',
                    'computed',
                    'watch',
                    'LIFECYCLE_HOOKS',
                    'methods',
                    ['template', 'render'],
                    'renderError',
                ],
            },
        ],
        'vue/require-prop-types': ['error'],
        'vue/require-valid-default-prop': ['error'],
    },
};