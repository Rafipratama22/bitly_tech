import React, { useEffect, useState } from "react";
import Swal from "sweetalert2";
import { useDispatch, useSelector } from "react-redux";
import { useHistory } from "react-router-dom";
import { loginUser } from "../actions/actions/Action";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

function Login() {
  const history = useHistory();
  const dispatch = useDispatch();
  const isLogin = useSelector((state) => state.user.isLogin)
  const [loggedCurrentUser, setCurrentUser] = useState({
    find: "",
    password: "",
  });

  const handleLogin = (e) => {
    e.preventDefault();
    dispatch(loginUser(loggedCurrentUser));
    console.log(isLogin, 'ini di handle')
  };

  // console.log(isLogin, 'ini di luar')
  useEffect(() => {
    if (isLogin || localStorage.getItem('access_token')) {
      Swal.fire({
        icon: "success",
        title: "Login success"
      });
      history.push("/");
    }
  }, [isLogin])


  return (
    <div className="login">
    <Form onSubmit={handleLogin}>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Email address</Form.Label>
        <Form.Control type="email" placeholder="Enter email" 
        onChange={(e) => 
            setCurrentUser({ ...loggedCurrentUser, find: e.target.value})
        }
        />
        <Form.Text className="text-muted">
        </Form.Text>
      </Form.Group>

      <Form.Group className="mb-3" controlId="formBasicPassword">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" 
        onChange={(e) => 
            setCurrentUser({...loggedCurrentUser, password: e.target.value})
        }
        />
      </Form.Group>
      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
    </div>
  )
}

export default Login;