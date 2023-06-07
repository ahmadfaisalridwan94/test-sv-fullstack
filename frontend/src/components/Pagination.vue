<template>
    <div class="flex items-center py-8 mx-auto">
        <a v-if="currentPage > 1" href="#" @click="goToPage(currentPage - 1)"
            class="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center mr-3">
            <i class="fas fa-arrow-left mr-2"></i> Previous
        </a>
        <template v-for="(page, index) in displayedPages">
            <a v-if="page === '...'" :key="'ellipsis-' + index" href="#"
                class="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center mx-1">
                {{ page }}
            </a>
            <a v-else :key="page" href="#" @click="goToPage(page)"
                :class="['h-10 w-10 font-semibold text-sm flex items-center justify-center', currentPage === page ? 'bg-blue-800 text-white' : 'text-gray-800 hover:bg-blue-600 hover:text-white']">
                {{ page }}
            </a>
        </template>
        <a v-if="currentPage < totalPages" href="#" @click="goToPage(currentPage + 1)"
            class="h-10 w-10 font-semibold text-gray-800 hover:text-gray-900 text-sm flex items-center justify-center ml-3">
            Next <i class="fas fa-arrow-right ml-2"></i>
        </a>
    </div>
</template>
  
  
<script>
export default {
    props: {
        currentPage: {
            type: Number,
            required: true
        },
        totalPages: {
            type: Number,
            required: true
        }
    },
    computed: {
        displayedPages() {
            const totalDisplayPages = 7; // Number of pages to display (excluding ellipsis)
            const pageNeighbours = Math.floor(totalDisplayPages / 2);
            const totalPages = this.totalPages;
            const currentPage = this.currentPage;

            if (totalPages <= totalDisplayPages) {
                return Array.from({ length: totalPages }, (_, i) => i + 1);
            }

            const pages = [];

            // Determine the range of pages to display
            let startPage = Math.max(2, currentPage - pageNeighbours);
            let endPage = Math.min(totalPages - 1, currentPage + pageNeighbours);

            // Add the first page
            pages.push(1);

            if (currentPage - pageNeighbours > 2) {
                pages.push('...');
            }

            // Add the intermediate pages
            for (let i = startPage; i <= endPage; i++) {
                pages.push(i);
            }

            if (currentPage + pageNeighbours < totalPages - 1) {
                pages.push('...');
            }

            // Add the last page
            pages.push(totalPages);

            return pages;
        }
    },
    methods: {
        goToPage(page) {
            // Emit an event to update the current page in the parent component
            this.$emit('update-page', page);          
        }
    }
};
</script>
  