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
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const searchResults = ref([]);

    const searchUSA = async () => {
      try {
        const response = await axios.get('http://localhost:3000/api/v1/search/usa');
        searchResults.value = response.data.hits.hits;
        console.log(searchResults.value[0]._index);
      } catch (error) {
        console.error('Error fetching search results:', error);
      }
    };

    onMounted(searchUSA);

    return { searchResults };
  },
};
</script>
