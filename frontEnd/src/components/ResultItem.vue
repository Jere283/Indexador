<template>
  <div>
    <h1>Search Results</h1>
    <div v-for="result in searchResults" :key="result._id">
      <div>
        <strong>ID:</strong> {{ result._id }}
      </div>
      <div>
        <strong>Score:</strong> {{ result._score }}
      </div>
      <div>
        <strong>Source:</strong> {{ JSON.stringify(result._source, null, 2) }}
      </div>
      <!-- Add more fields as needed -->
    </div>
  </div>
</template>

<script>
import { ref, watch } from 'vue';
import axios from 'axios';

export default {
  props: ['searchTerm'],
  setup(props) {
    const searchResults = ref([]);

    const search = async (term) => {
      try {
        const response = await axios.get(`http://localhost:3000/api/v1/search/${term}`);
        searchResults.value = response.data.hits.hits;
      } catch (error) {
        console.error('Error fetching search results:', error);
      }
    };

    watch(() => props.searchTerm, (newVal) => {
      if (newVal && newVal.length > 0) {
        search(newVal);
      }
    });

    return { searchResults };
  },
};
</script>
