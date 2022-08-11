import React from "react";
import { useParams } from "react-router-dom";
import {useEffect} from 'react'

function Bitly() {
  const { id } = useParams();
  useEffect(() => {
    
  })
  return (
    <div className="home">
      <h1>Bitly</h1>
    </div>
  );
}

export default Bitly;
