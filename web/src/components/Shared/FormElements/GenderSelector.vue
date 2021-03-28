<template>
  <label class="block">
    <span class="text-gray-700">{{ text }}</span>
    <div class="mt-2">
      <div>
        <label class="inline-flex items-center">
          <input
            name="gender"
            value="1"
            type="radio"
            v-model="genderlocal"
            class="rounded border-gray-300 text-indigo-600 shadow-sm focus:border-indigo-300 focus:ring focus:ring-offset-0 focus:ring-indigo-200 focus:ring-opacity-50"
          />
          <span class="ml-2">Male</span>
        </label>
        <br />
        <label class="inline-flex items-center">
          <input
            name="gender"
            value="2"
            type="radio"
            v-model="genderlocal"
            class="rounded border-gray-300 text-indigo-600 shadow-sm focus:border-indigo-300 focus:ring focus:ring-offset-0 focus:ring-indigo-200 focus:ring-opacity-50"
          />
          <span class="ml-2">Female</span>
        </label>
      </div>
    </div>
  </label>
</template>
<script>
import { computed } from 'vue'
export default {
  name: 'gender-selector',
  props: {
    gender: Number,
    text: String,
  },
  setup(props, { emit }) {
    const genderlocal = computed({
      get() {
        return props.gender
      },
      set(newvalue) {
        // Radio buttons values can't map to numbers
        emit('update:gender', parseInt(newvalue))
      },
    })
    return {
      genderlocal,
    }
  },
}
</script>
