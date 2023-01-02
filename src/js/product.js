function createProduct(event){
  event.preventDefault();
  const productData = new FormData(document.getElementById('product-form'))
  fetch('http://localhost:3000/api/v1/products/', {
    method: 'POST',
    body: JSON.stringify({
      Name: productData.get('Name'),
      Price: parseFloat(productData.get('Price')),
    }),
  })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      window.location.href = '/';
    }
    )
}

function editProduct(event){
  event.preventDefault();
  const productData = new FormData(document.getElementById('product-form'))
  fetch(`http://localhost:3000/api/v1/products/${productData.get('ProductID')}`, {
    method: 'PATCH',
    body: JSON.stringify({
      Name: productData.get('Name'),
      Price: parseFloat(productData.get('Price')),
    }),
  })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      window.location.href = '/';
    }
    )
    .catch((error) => {
      console.error('Error:', error);
    }
    );
}

const deleteProduct = async (productId) => {
  try {
    const result = await fetch(`http://localhost:3000/api/v1/products/${productId}`, {
      method: 'DELETE'
    });
    await result.json();
    document.getElementById(`product-${productId}`).remove();
  } catch (error) {
    console.error(error)
  }
}
