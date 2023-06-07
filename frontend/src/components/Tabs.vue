<template>
    <div class="border-b border-gray-200 dark:border-gray-700">
        <ul class="flex flex-wrap -mb-px text-sm font-medium text-center text-gray-500 dark:text-gray-400">
            <li class="mr-2">
                <a href="#" :class="{ active: activeTab === 'publish' }" @click="activateTab('publish');"
                    class="inline-flex p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300 group">
                    Published
                </a>
            </li>
            <li class="mr-2">
                <a href="#" :class="{ active: activeTab === 'draft' }" @click="activateTab('draft');"
                    class="inline-flex p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300 group">
                    Drafts
                </a>
            </li>
            <li class="mr-2">
                <a href="#" :class="{ active: activeTab === 'trash' }" @click="activateTab('trash');"
                    class="inline-flex p-4 border-b-2 border-transparent rounded-t-lg hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300 group">
                    Trashed
                </a>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            activeTab: "publish", // Set the default active tab
        };
    },
    methods: {
        activateTab(tab) {
            this.activeTab = tab; // Set the active tab to the clicked tab
            this.$emit('update-status', tab); // Emit the status to the parent component
        },
    },
};
</script>

<style scoped>
.active {
    /* Define styles for the active tab */
    /* For example, you can set the text color or border color */
    color: blue;
    border-color: blue;
}
</style>
