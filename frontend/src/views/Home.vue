<template>
	<div class="max-w-6xl mx-auto py-8">
		<button @click="goToAddNew" type="button"
			class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">
			Add New
		</button>
		<button @click="goToPreview" type="button"
			class="text-blue-700 hover:text-white border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:border-blue-500 dark:text-blue-500 dark:hover:text-white dark:hover:bg-blue-500 dark:focus:ring-blue-800">
			Preview
		</button>

		<Tabs @update-status="updatePostStatus" />

		<TableComponent :posts="posts" @update-status="updatePostStatus" :activeTab="this.activeTab" />

		<Pagination :current-page="currentPage"  :total-pages="totalPages" @update-page="updateCurrentPage" />

	</div>
</template>

<script>

import Tabs from '../components/Tabs.vue';
import TableComponent from '../components/TableComponent.vue';
import Pagination from '../components/Pagination.vue';
import axios from 'axios'

export default {
	components: {
		TableComponent,
		Tabs,
		Pagination
	},

	data() {
		return {
			posts: [],
			currentPage: 1,
			totalPages: 1,
			perPage: 5,
			activeTab: 'publish',
		};
	},
	mounted() {
		this.fetchPosts('publish');
	},
	methods: {
		updatePostStatus(status) {
			this.currentPage = 1
			this.activeTab = status
			this.fetchPosts(status); // Call the fetchPosts method with the updated status
		},
		goToPreview() {
			this.$router.push('/preview');
		},
		goToAddNew() {
			this.$router.push('/post/create');
		},
		fetchPosts(status) {
			
			axios.get(`/count_article/${status}`)
				.then(response => {
					var countData = response.data.Data;
					this.totalPages = Math.ceil(countData / this.perPage);
					this.offset = (this.currentPage - 1) * this.perPage;

					axios.get(`/article/${this.perPage}/${this.offset}?status=${status}`)
						.then(response => {
							this.posts = response.data;
						})
						.catch(error => {
							console.error(error);
						});
				})
				.catch(error => {
					console.error(error);
				});


		},
		updateCurrentPage(page) {
			this.currentPage = page;
			// Fetch posts for the updated page
			this.fetchPosts(this.activeTab);
		},
	},
}; 
</script>
