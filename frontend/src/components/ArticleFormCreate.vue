<template>
    <form @submit.prevent="handleSubmit">
        <div class="mb-6">
            <label for="title" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Title</label>
            <input type="text" id="title" v-model="title" @change="handleChangeTitle" @keyup="handleKeyupTitle"
                :class="{ 'bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500': true, 'border-red-500': titleInvalid }"
                placeholder="Post Title" required>
            <span v-for="errMsg in errorsTitle" class="text-red-600">
                {{ errMsg }}
            </span>

        </div>
        <div class="mb-6">
            <label for="content" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Content</label>
            <textarea id="content" rows="10" v-model="content" @change="handleChangeContent" @keyup="handleKeyupContent"
                :class="{ 'bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500': true, 'border-red-500': contentInvalid }"
                required></textarea>
            <span v-for="errMsg in errorsContent" class="text-red-600">
                {{ errMsg }}
            </span>
        </div>
        <div class="mb-6">
            <label for="category" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Category</label>
            <input type="text" id="category" v-model="category" @change="handleChangeCategory" @keyup="handleKeyupCategory"
                :class="{ 'bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500': true, 'border-red-500': categoryInvalid }"
                placeholder="Post Category" required>
            <span v-for="errMsg in errorsCategory" class="text-red-600">
                {{ errMsg }}
            </span>
        </div>

        <div class="flex space-x-2">
            <button type="button" @click="publishPost"
            class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            Publish</button>
        <button type="button" @click="savePostDraft"
            class="text-white bg-yellow-400 hover:bg-yellow-500 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            Draft</button>
        <button type="button" @click="goHome"
            class="text-white bg-blue-400 hover:bg-blue-500 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            Back</button>
        </div>
    </form>
</template>

<script>

import axios from "axios";
import { toast } from 'vue3-toastify';

import 'vue3-toastify/dist/index.css';


export default {
    setup() {

    },
    data() {
        return {
            title: '',
            content: '',
            category: '',
            errorsTitle: [],
            errorsContent: [],
            errorsCategory: [],
            titleInvalid: false,
            contentInvalid: false,
            categoryInvalid: false,
        }
    },
    methods: {

        goHome() {
            this.$router.push('/');
        },

        async publishPost() {

            await axios
                .post("article/", JSON.stringify({
                    title: this.title,
                    content: this.content,
                    category: this.category,
                    status: 'publish'
                }))
                .then((response) => {
                    toast.success("Published !", {
                        autoClose: 1000,
                    });
                })
                .catch(error => {
                    if (error.response) {
                        
                        var errData = error.response.data
                        this.handleValidation(errData)

                    } else {
                        // Something happened in setting up the request that triggered an error
                        //console.log('Error', error.message);
                    }
                });
        },

        async savePostDraft() {

            await axios
                .post("article/", JSON.stringify({
                    title: this.title,
                    content: this.content,
                    category: this.category,
                    status: 'draft'
                }))
                .then((response) => {
                    toast.success("Saved to draft !", {
                        autoClose: 1000,
                    });
                })
                .catch(error => {
                    if (error.response) {

                        var errData = error.response.data
                        this.handleValidation(errData)                        

                    } else {
                        // Something happened in setting up the request that triggered an error
                        console.log('Error', error.Title);
                        this.errorsTitle = error.Title
                    }
                });

        },

        handleValidation(errData) {

            if (errData.hasOwnProperty("Title")) {

                this.errorsTitle = errData.Title
                if ((errData.Title).length > 0) {
                    this.titleInvalid = true
                } else {
                    this.titleInvalid = false
                }

            }

            if (errData.hasOwnProperty("Content")) {
                this.errorsContent = errData.Content
                if ((errData.Content).length > 0) {
                    this.contentInvalid = true
                } else {
                    this.contentInvalid = false
                }
            }

            if (errData.hasOwnProperty("Category")) {
                this.errorsCategory = errData.Category
                if ((errData.Category).length > 0) {
                    this.categoryInvalid = true
                } else {
                    this.categoryInvalid = false
                }
            }

        },

        handleChangeTitle() {
            this.titleInvalid = false
        },

        handleChangeContent() {
            this.contentInvalid = false
        },

        handleChangeCategory() {
            this.categoryInvalid = false
        },

        handleKeyupTitle() {
            this.errorsTitle = []
        },

        handleKeyupContent() {
            this.errorsContent = []
        },

        handleKeyupCategory() {
            this.errorsCategory = []
        }
    },
};
</script>
