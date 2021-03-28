<template>
  <div v-show="editorOpen" class="relative p-3 overflow-hidden">
    <div
      class="relative p-4 bg-gray-100 dark:bg-gray-700 shadow-lg rounded-3xl border-green-800 border-2"
    >
      <div class="grid grid-cols-1 gap-6">
        <label class="block">
          <span class="text-gray-700">Full name</span>
          <input
            type="text"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 dark:bg-gray-700"
            placeholder="Name"
            v-model="editingPerson.fullname"
          />
        </label>
        <label class="block">
          <span class="text-gray-700">Birthday</span>
          <input
            type="date"
            v-model="editingPerson.birthday"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
          />
        </label>
        <gender-selector v-model:gender="editingPerson.gender" />
        <label class="block">
          <span class="text-gray-700">Gender</span>
          <div class="mt-2">
            <div>
              <label class="inline-flex items-center">
                <input
                  name="gender"
                  value="1"
                  type="radio"
                  v-model.number="editingPerson.gender"
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
                  v-model.number="editingPerson.gender"
                  class="rounded border-gray-300 text-indigo-600 shadow-sm focus:border-indigo-300 focus:ring focus:ring-offset-0 focus:ring-indigo-200 focus:ring-opacity-50"
                />
                <span class="ml-2">Female</span>
              </label>
            </div>
          </div>
        </label>
        <div class="block">
          <standard-button
            type="submit"
            class="flex-1"
            text="Submit"
            @click="closeEditor"
          />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import StandardButton from '../Shared/Buttons/StandardButton.vue'
import GenderSelector from '../Shared/FormElements/GenderSelector.vue'

export default {
  name: 'person-editor',
  components: { StandardButton, GenderSelector },
  props: {
    person: Object,
  },
  setup(props) {
    const onSubmit = () => {
      return props.person
    }

    return {
      onSubmit,
    }
  },
}
</script>
