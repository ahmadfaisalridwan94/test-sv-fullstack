<script>

import { TrashIcon } from "@heroicons/vue/24/solid";
import { PencilSquareIcon } from "@heroicons/vue/24/outline";
import axios from "axios"

export default {
	components: {
		TrashIcon,
		PencilSquareIcon
	},
	data() {
		return {
			tableHeaderItems: ["Title", "Category", "Action"],	
			activeTab: ''		
		};
	},
	methods: {
		activateTab(tab) {
			this.activeTab = tab; // Set the active tab to the clicked tab
		},
		handleEdit(id) {
			this.$router.push(`/post/edit/${id}`);			
		},
		showAlertConfirm(id, paramActiveTab) {
			this.$swal(
				{
					title: 'Are you sure?',
					text: "You won't be able to revert this!",
					icon: 'warning',
					showCancelButton: true,
					confirmButtonColor: '#3085d6',
					cancelButtonColor: '#d33',
					confirmButtonText: 'Yes, delete it!'
				}
			).then((result) => {
				if (result.isConfirmed) {
					//const config = { headers: { 'Content-Type': 'application/json'} };
					axios
						.post(`article/${id}/`)
						.then((response) => {
							this.$swal(
								{
									text: "Post moved to trash!",
									icon: 'success',
								});
														
							this.$emit('update-status', paramActiveTab);
						})
						.catch(error => {
							if (error.response) {
								// The request was made, but the server responded with a status code outside the 2xx range
								console.log(error.response.data);
								console.log(error.response.status);
								console.log(error.response.headers);
							} else {
								// Something happened in setting up the request that triggered an error
								console.log('Error', error.message);
							}
						});

				}
			});
		}
	},
	props: {
		posts: {
			type: Array,
			required: true,
		},
		activeTab:{
			type: String,
			required: true 
		}
	},
};
</script>

<template>
	<div class="relative overflow-x-auto mt-3">
		<table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
			<thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
				<tr>
					<th v-for="hItem in tableHeaderItems" scope="col" class="px-6 py-3">
						{{ hItem }}
					</th>
				</tr>
			</thead>
			<tbody>
				<tr v-for="post in posts" :key="post.id" :post="post"
					class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
					<td scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
						{{ post.Title }}
					</td>
					<td class="px-6 py-4">{{ post.Category }}</td>
					<td class="px-6 py-4">

						<div class="flex space-x-1" v-if="post.Status != 'trash'">
							<button title="Edit" @click="handleEdit(post.Id)"
								class="bg-grey-light hover:bg-blue text-blue-darkest font-bold py-2 px-1 rounded inline-flex items-center">
								<PencilSquareIcon class="h-6 w-6 text-blue-500" />
							</button>
							<button title="Delete" @click="showAlertConfirm(post.Id, activeTab)"
								class="bg-grey-light hover:bg-blue text-blue-darkest font-bold py-2 px-1 rounded inline-flex items-center">
								<TrashIcon class="h-6 w-6 text-red-500" />
							</button>
						</div>

					</td>
				</tr>
			</tbody>
		</table>
	</div>
</template>
