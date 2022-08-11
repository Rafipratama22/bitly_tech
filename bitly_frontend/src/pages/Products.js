import React, {useEffect} from 'react';
import ListGroup from 'react-bootstrap/ListGroup';
import { useDispatch, useSelector } from "react-redux";
import { fetchProducts } from "../actions/actions/Action";

function Products() {
  const dispatch = useDispatch()
  const allProducts = useSelector((state) => state.user.allProducts)
  useEffect(() => {
    dispatch(fetchProducts())
  }, [])
  return (
    <div className='products'>
        <ListGroup>
          {
            allProducts.map(product => (
              <ListGroup.Item key={product.ID}>
                <p>Url: {product.ModifyUrl}</p>
                <p>Destination: {product.DestinationUrl}</p>
                <p>Title: {product.Title}</p>
                <p>Source: {product.Source}</p>
                <p>Medium: {product.Medium}</p>
                <p>Campaign: {product.Campaign}</p>
              </ListGroup.Item>
            )) 
          }
        </ListGroup>
    </div>
  );
}

export default Products;
