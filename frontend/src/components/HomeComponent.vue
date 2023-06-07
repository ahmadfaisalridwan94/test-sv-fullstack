<script setup>
import { ref } from "vue";

defineProps({
	msg: String,
});

const count = ref(0);
</script>

<template>
	<div class="flex items-center justify-center h-screen w-100">
		<div class="max-w-md w-full bg-white rounded-lg shadow-lg p-8">
			<h1 class="text-4xl font-semibold mb-6">Welcome to Our Website</h1>
			<p class="text-gray-700 mb-6">
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque maximus
				neque vitae magna ultrices, et mattis dolor convallis. Aliquam faucibus
				urna ac nulla bibendum, nec dictum arcu blandit. Curabitur auctor
				condimentum rutrum. Sed in neque risus.
			</p>
			<router-link
				:to="{ name: 'post' }"
				class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
			>
				Post
			</router-link>
		</div>
	</div>
</template>

<style scoped>
.read-the-docs {
	color: #888;
}
</style>
