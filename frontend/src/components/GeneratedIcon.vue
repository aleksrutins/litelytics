<script setup lang="ts">
    defineProps<{
        size: string
    }>()
    const gradientChoices = [
        'from-sky-500 to-indigo-500',
        'from-cyan-500 to-blue-500',
        'from-violet-500 to-fuschia-500',
        'from-purple-500 to-pink-500',
        'from-red-500 to-orange-500'
    ]
    const chosenGradient = gradientChoices[Math.round(Math.random() * (gradientChoices.length - 1))];
</script>
<template>
    <div :class="`bg-gradient-to-br ${chosenGradient} rounded-full`" :style="`width: ${size}; height: ${size};`"></div>
</template>