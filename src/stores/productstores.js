import { books } from '../Data_set/products';
import { defineStore } from 'pinia';

export const useProductStore = defineStore('productStore', {
  state: () => ({
    products: []
  }),

  actions: {
    fetchProducts() {
      this.products = books.map(product => ({
        ...product,
        image: new URL(`../assets/images/${product.image}`, import.meta.url).href
      }));
    }
  }
});
