<template>
  <div
    v-if="modelValue"
    class="fixed z-10 inset-0 overflow-y-auto bg-gray-500 bg-opacity-50"
    aria-labelledby="modal-title"
    role="dialog"
    aria-modal="true"
  >
    <!-- Modal content -->
    <div
      class="rounded-lg modal-content mx-4 sm:mx-8 lg:mx-auto max-w-screen-lg relative text-white"
    >
      <div class="mb-7">
        <button @click="closeModal" class="absolute right-6 cursor-pointer">
          X
        </button>
      </div>
      <!--full content -->
      <div>
        <h3 class="text-center font-bold text-2xl">Email Details</h3>
        <p><strong>From:</strong> {{ content.From }}</p>
        <p><strong>To:</strong> {{ content.To }}</p>
        <p><strong>Subject:</strong> {{ content.Subject }}</p>
        <p><strong>Body:</strong></p>
        <p
          class="leading-relaxed whitespace-pre-wrap whitespace-pre-wrap whitespace-moz-pre-wrap whitespace-pre-wrap whitespace-o-pre-wrap break-words"
        >
          {{ content.body }}
        </p>
      </div>
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
  background: linear-gradient(to bottom right, #282c35, #1a1d23);
  padding: 20px;
  margin-top: 15px;
}

p {
  margin-bottom: 4px !important;
  white-space: pre-wrap;
  white-space: -moz-pre-wrap;
  white-space: -o-pre-wrap;
  word-wrap: break-word;
}
</style>
