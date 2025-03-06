<template>
  <div class="flex justify-center items-center min-h bg-gray-100 p-10">
    <article class="bg-white shadow-lg rounded-lg p-6 w-full max-w-lg">

      <!-- Step Indicator -->
      <header class="flex justify-between mb-10">
        
        <div
          v-for="(step, index) in formGroup"
          :key="'step' + index"
          class="w-10 h-10 flex items-center justify-center rounded-full text-white font-bold transition-all"
          :class="index === formPosition ? 'bg-blue-500 scale-110' : 'bg-gray-400'"
        >
          {{ index + 1 }}
        </div>
      </header>

      <!-- Form Section -->
      <section :class="animation" class="transition-all duration-500">
        <div class=" justify-items-center ">       
           <img src="A:/Project/My_Project/book_family/src/image/profile-1341-svgrepo-com.svg" class="rounded-full w-20 h-20" alt="">
        </div>
        <h2 class="text-xl font-semibold text-gray-700 text-center mb-4">
          {{ formGroup[formPosition].title }}
        </h2>

        <div class="space-y-4">
          <div v-for="(field, index) in formGroup[formPosition].fields" :key="'field' + index">
            <label class="block text-gray-600 mb-1">{{ field.label }}</label>
            <input
              type="text"
              v-model="field.value"
              required
              class="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-blue-400 focus:border-blue-400 transition"
            />
          </div>
        </div>

        <!-- Buttons -->
        <div class="flex justify-between mt-6">
          <button
            v-if="formPosition > 0"
            @click="prevStep"
            class="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition"
          >
            Back
          </button>
          <button
            v-if="formPosition + 1 < formGroup.length"
            @click="nextStep"
            class="ml-auto px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
          >
            Next
          </button>
          <button
            v-if="formPosition + 1 === formGroup.length"
            @click="submitForm"
            class="ml-auto px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition"
          >
            Submit
          </button>
        </div>
      </section>
    </article>
  </div>
</template>
<script>
import axios from "axios"; // Import Axios

export default {
  data() {
    return {
      formPosition: 0,
      animation: "fade-in",
      formGroup: [
        {
          title: "Personal Details",
          fields: [
            { label: "First Name", value: "" },
            { label: "Second Name", value: "" },
            { label: "Age", value: "" },
          ],
        },
        {
          title: "Address",
          fields: [
            { label: "City", value: "" },
            { label: "Zip Code", value: "" },
            { label: "County", value: "" },
            { label: "State", value: "" },
          ],
        },
        {
          title: "Academic Details",
          fields: [
            { label: "Academic qualification", value: "" },
            { label: "College Attended", value: "" },
            { label: "Field Of Study", value: "" },
          ],
        },
      ],
    };
  },
  methods: {
    nextStep() {
      this.animation = "fade-out";
      setTimeout(() => {
        this.formPosition += 1;
        this.animation = "fade-in";
      }, 300);
    },
    prevStep() {
      this.animation = "fade-out";
      setTimeout(() => {
        this.formPosition -= 1;
        this.animation = "fade-in";
      }, 300);
    },
    async submitForm() {
      const isValid = this.validateForm();
      if (isValid) {
        // Prepare the data to be sent to the backend
        const userData = this.getFormData();

        try {
          const response = await axios.post("http://localhost:9010/user/", userData);
          // Handle success response
          if (response.status === 200) {
            alert("User registered successfully!");
          }
        } catch (error) {
          // Handle error response
          alert("Error: " + (error.response?.data?.message || "Something went wrong!"));
        }
      } else {
        alert("Please fill out all fields before submitting.");
      }
    },
    validateForm() {
      for (let step of this.formGroup) {
        for (let field of step.fields) {
          if (!field.value.trim()) {
            return false;
          }
        }
      }
      return true;
    },
    getFormData() {
      const formData = {};
      this.formGroup.forEach(step => {
        step.fields.forEach(field => {
          // Use the field label as the key, and the value entered as the value
          formData[field.label] = field.value;
        });
      });
      return formData;
    }
  },
};
</script>
