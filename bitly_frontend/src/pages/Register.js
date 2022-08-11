import React, { useEffect, useState } from "react";
import Swal from "sweetalert2";
import { useDispatch, useSelector } from "react-redux";
import { useHistory } from "react-router-dom";
import { regisUser } from "../actions/actions/Action";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

function Register() {
  const history = useHistory();
  const dispatch = useDispatch();
  const isRegis = useSelector((state) => state.user.isRegis)
  const [loggedCurrentUser, setCurrentUser] = useState({
    email: "",
    password: "",
    username: ""
  });

  const handleRegister = (e) => {
    e.preventDefault();
    dispatch(regisUser(loggedCurrentUser));
    console.log(isRegis, 'ini di handle')
  };

  // console.log(isLogin, 'ini di luar')
  useEffect(() => {
    if (isRegis || localStorage.getItem('access_token')) {
      Swal.fire({
        icon: "success",
        title: "Register success"
      });
      history.push("/login");
    }
  }, [isRegis])


  return (
    <div className="register">
    <Form onSubmit={handleRegister}>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Email address</Form.Label>
        <Form.Control type="email" placeholder="Enter email"
        onChange={(e) => 
            setCurrentUser({ ...loggedCurrentUser, email: e.target.value})
        }
        />
        <Form.Text className="text-muted">
        </Form.Text>
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Label>Username</Form.Label>
        <Form.Control type="text" placeholder="Enter username"
        onChange={(e) => 
            setCurrentUser({ ...loggedCurrentUser, username: e.target.value})
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

export default Register;