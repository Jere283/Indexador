<template>

<div class="container" v-for="result in searchResults" :key="result._id">
  <div class="min-h-screen flex items-center justify-center px- mt-4 mb-4">
    <div class="max-w-4xl bg-white w-full rounded-lg shadow-xl">
      <div class="p-4 border-b">
        <h2 class="text-2xl">Email Information</h2>
        <p class="text-sm text-gray-500">Message Id: {{ result._source.Message_id }} </p>
      </div>
      <div>
        <div
          class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
        >
          <p class="text-gray-600">Subject</p>
          <p>{{ result._source.Subject }}</p>
        </div>
        <div
          class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
        >
          <p class="text-gray-600">Date</p>
          <p>{{ result._source.Date }}</p>
        </div>
        <div
          class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
        >
          <p class="text-gray-600">From</p>
          <p>{{ result._source.From }}</p>
        </div>
        <div
          class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
        >
          <p class="text-gray-600">To</p>
          <p>{{ result._source.To }}</p>
        </div>
        <div
          class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
        >
          <p class="text-gray-600">Body</p>
          <p>
            {{ result._source.body }}
          </p>
        </div>
        
      </div>
    </div>
  </div>
</div>
  
  
</template>

<script>
import { ref, watch } from "vue";
import axios from "axios";

export default {
  props: ["searchTerm"],
  setup(props) {
    const searchResults = ref([]);

    const search = async (term) => {
      try {
        const response = await axios.get(
          `http://localhost:3000/api/v1/search/${term}`
        );
        searchResults.value = response.data.hits.hits;

        console.log(searchResults.value[0]._source);
      } catch (error) {
        console.error("Error fetching search results:", error);
      }
    };

    watch(
      () => props.searchTerm,
      (newVal) => {
        if (newVal && newVal.length > 0) {
          search(newVal);
        }
      }
    );

    return { searchResults };
  },
};

var i = 0;
</script>

<style>
.container {
  max-width: 100%;  /* Ensure the container does not exceed the viewport width */
  overflow: hidden; /* Hide any content that overflows the container */
  word-wrap: break-word; /* Break long words and prevent overflow */
  /* Add additional styles as needed */
}

/* Add any other styles for your page layout */
</style>