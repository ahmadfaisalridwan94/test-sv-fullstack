<template>
    <div class="max-w-3xl mx-auto py-3">

        <button @click="goToHome" class="text-blue-700 hover:text-white border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:border-blue-500 dark:text-blue-500 dark:hover:text-white dark:hover:bg-blue-500 dark:focus:ring-blue-800">Posts</button>

        <h2 class="text-4xl mb-4 text-center font-extrabold dark:text-white">Blog Posts</h2>

        <PostBox v-for="post in posts" :key="post.id" :post="post" />

        <Pagination :current-page="currentPage" :total-pages="totalPages" @update-page="updateCurrentPage" />

    </div>
</template>

<script>
import PostBox from '../components/PostBox.vue'
import Pagination from '../components/Pagination.vue'
import axios from 'axios'

export default {
    components: {
        PostBox,
        Pagination,
    },
    data() {
        return {
            posts: [],
            currentPage: 1,
            totalPages: 1,
            perPage:5,
            offset: 0
        };
    },
    mounted() {
        this.fetchPosts();
    },
    methods: {
        fetchPosts() {

            axios
				.get(`/count_article/publish`)
				.then(response => {
					var countData = response.data.Data;
					this.totalPages = Math.ceil(countData / this.perPage);
					this.offset = (this.currentPage - 1) * this.perPage;

                    axios
                        .get(`/article/${this.perPage}/${this.offset}?status=publish`)
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
            this.offset = this.currentPage - 1
			// Fetch posts for the updated page
			this.fetchPosts();
		},
        goToHome() {
			this.$router.push('/');
		},
    },
};
</script>