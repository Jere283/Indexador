<template>
  <div class="mt-5">
    <div class="px-4 sm:px-8 max-w-5xl m-auto">
      <h1 class="text-white text-center font-semibold text-sm">
        Search results: {{ searchResults.length }}
      </h1>
      <ul class="border border-transparent rounded overflow-hidden shadow-md">
        <li
          v-for="result in searchResults"
          :key="result._id"
          @click="openModal(result)"
          class="mt-4 px-4 py-2 bg-blue-gray-800 hover:bg-sky-50 hover:text-sky-900 border-b last:border-none border-transparent rounded transition-all duration-300 ease-in-out"
        >
          <strong>From:</strong> {{ result._source.From }} <strong>To:</strong>
          {{ result._source.To }} <strong>Subject:</strong>
          {{ result._source.Subject }}
        </li>
      </ul>
    </div>
    <Modal
      :modelValue="showModal"
      @update:modelValue="(val) => (showModal = val)"
      :content="modalContent"
    />
  </div>
</template>

<script>
import { ref, watch } from "vue";
import axios from "axios";
import Modal from "./Modal.vue";

export default {
  components: {
    Modal,
  },
  props: ["searchTerm"],
  setup(props) {
    const searchResults = ref([]);
    const showModal = ref(false);
    const modalContent = ref("");

    const search = async (term) => {
      try {
        const response = await axios.get(
          `http://localhost:3000/api/v1/search/${term}`
        );
        searchResults.value = response.data.hits.hits;
      } catch (error) {
        console.error("Error fetching search results:", error);
      }
    };

    const openModal = (result) => {
      console.log(result);
      showModal.value = true;
      modalContent.value = result._source;
    };

    watch(
      () => props.searchTerm,
      (newVal) => {
        if (newVal && newVal.length > 0) {
          search(newVal);
        }
      }
    );

    return { searchResults, showModal, modalContent, openModal };
  },
};
</script>

<style>
li {
  background-color: #012143 !important;
  color: #ffffff;
}
</style>
