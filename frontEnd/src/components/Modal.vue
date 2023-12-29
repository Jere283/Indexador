<template>
  <div
    v-if="modelValue"
    class="fixed z-10 inset-0 overflow-y-auto bg-gray-500 bg-opacity-50"
    aria-labelledby="modal-title"
    role="dialog"
    aria-modal="true"
  >
    <!-- Modal content -->
    <div class="modal-content">
      <button @click="closeModal" class="close-button">X</button>
      <!--full content -->
      <p class="modal-text">{{ content }}</p>
    </div>
  </div>
</template>

<script>
export default {
  emits: ["update:modelValue"],
  props: ["modelValue", "content"],
  methods: {
    closeModal() {
      this.$emit("update:modelValue", false);
    },
  },
};
</script>
<style>
.modal-content {
  background-color: white;
  padding: 20px;
  margin: 50px auto;
  max-width: 500px;
  position: relative;
}

.close-button {
  position: absolute;
  right: 10px;
  top: 10px;
  cursor: pointer;
}

.modal-text {
  white-space: pre-wrap; /* CSS3 */
  white-space: -moz-pre-wrap; /* Firefox */
  white-space: -pre-wrap; /* Opera <7 */
  white-space: -o-pre-wrap; /* Opera 7 */
  word-wrap: break-word; /* IE */
}
</style>
