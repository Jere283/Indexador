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
      <button
        @click="closeModal"
        class="absolute top-10 right-10 cursor-pointer"
      >
        X
      </button>
      <!--full content -->
      <p
        class="whitespace-pre-wrap whitespace-pre-wrap whitespace-moz-pre-wrap whitespace-pre-wrap whitespace-o-pre-wrap break-words"
      >
        {{ content }}
      </p>
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
</style>
