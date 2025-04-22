<script>

defineProps({
  vacancy: {
    type: IVacancy,
    required: true
  }
})
defineEmits(['apply']);
</script>

<template>
  <div class="border-1 border-gray-300 rounded-lg p-4 mb-4 shadow-md hover:shadow-lg transition-shadow duration-300">
    <h2 class="font-medium text-4">{{ vacancy.title }}</h2>
    <p class="text-2">{{ vacancy.description }}</p>
    <div class="flex justify-between items-center mt-4">
            <p class="text-2">Location: {{ vacancy.location }}</p>
    <p>Salary: {{ vacancy.salary }}</p>
    </div>
    <button @click="apply">Apply Now</button>
  </div> 
</template>       