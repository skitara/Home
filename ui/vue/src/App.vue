<script setup>
import { ref } from 'vue';

// Define variables for saving data of form
const title = ref('');
const content = ref('');

// Function to submit data to server
const submitForm = async () => {
  const postData = {
    title: title.value,
    content: content.value,
  };

  // Send POST to server
  const response = await fetch('http://localhost:8080/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(postData),
  });

  if (response.ok) {
    alert('Post created successfully!');
    // Clean form after successful creation
    title.value = '';
    content.value = '';
  } else {
    alert('Error creating post!');
  }
};
</script>

<template>
  <div>
    <h1>Blog</h1>
    
    <!-- Form -->
    <form @submit.prevent="submitForm">
      <div>
        <label for="title">Title:</label>
        <input type="text" id="title" v-model="title" placeholder="Enter post title" required />
      </div>
      <div>
        <label for="content">Content:</label>
        <textarea id="content" v-model="content" placeholder="Enter post content" required></textarea>
      </div>
      <button type="submit">Create Post</button>
    </form>
  </div>
</template>

<style scoped>
/* Simple stylization for form */
form {
  display: flex;
  flex-direction: column;
  max-width: 400px;
  margin: 20px auto;
}

label {
  font-weight: bold;
}

input, textarea {
  padding: 10px;
  margin: 5px 0 15px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button {
  padding: 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}
</style>
