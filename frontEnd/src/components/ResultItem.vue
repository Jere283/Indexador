<template>

  <div class="mt-5" >
    <div class="px-4 sm:px-8 max-w-5xl m-auto">
        <h1 class="text-center font-semibold text-sm">Search results: {{ searchResults.length }}</h1>
        <ul class="border border-transparent rounded overflow-hidden shadow-md">
            <li  v-for="result in searchResults" :key="result._id" 
            class=" mt-4 px-4 py-2 bg-white hover:bg-sky-100 hover:text-sky-900 border-b last:border-none border-transparent rounded transition-all duration-300 ease-in-out">
            <strong>From:</strong> {{ result._source.From  }}  <strong>To:</strong> {{ result._source.To }}  <strong>Subject:</strong> {{ result._source.Subject }} </li>
        </ul>
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

        console.log(searchResults.value[5]._source);
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

</script>