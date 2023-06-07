<script>

export default {
	props: {
		post: {
			type: Object,
			required: true,
		},
	},
};

</script>

<template>
	<div class="bg-white rounded-lg shadow-lg p-6 mt-2">
		<div class="mb-6">
			<h2 class="text-2xl font-semibold mb-2">{{ post.Title }}</h2>
			<p class="text-gray-700">
				{{ post.Content }}
			</p>
		</div>
	</div>
</template>

<style scoped>
.read-the-docs {
	color: #888;
}
</style>
