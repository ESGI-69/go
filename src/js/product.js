const deleteProduct = async (productId) => {
  try {
    const result = await fetch(`http://localhost:3000/api/products/${productId}`, {
      method: 'DELETE'
    });
    await result.json();
    document.getElementById(`product-${productId}`).remove();
  } catch (error) {
    console.error(error)
  }
}
