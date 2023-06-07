<template>
	<Navbar />
	<main class="bg-gray-100">
		<div class="max-w-[1200px] mx-auto">
			<router-view />
		</div>
	</main>
</template>

<script setup>
import Navbar from "../components/Navbar.vue";
</script>
