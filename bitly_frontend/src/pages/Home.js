import React, {useState, useEffect} from "react";
import { useDispatch, useSelector } from "react-redux";
import { makeUrl } from "../actions/actions/Action";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';


function Home() {
  const dispatch = useDispatch();
  const isUrl = useSelector((state) => state.user.isUrl)
  const [log, setLog] = useState({
    destination_url: ""
  });

  const handleUrl = (e) => {
    e.preventDefault();
    dispatch(makeUrl(log))
  }

  useEffect(() => {
    console.log(isUrl)
  }, [isUrl])

  return (
    <div className="home">
      <Form onSubmit={handleUrl}>
        <Form.Group className="mb-3">
          <Form.Label>Enter Url</Form.Label>
          <Form.Control
            type="text"
            placeholder="Enter url"
            onChange={(e) =>
              setLog({ ...log, destination_url: e.target.value })
            }
          />
          {
            isUrl 
            ? <Form.Text className="muted" id="passwordHelpBlock">{isUrl}</Form.Text>
            : <Form.Text className="muted" hidden >{isUrl}</Form.Text>
          }
        </Form.Group>

        <Button variant="primary" type="submit">
          Submit
        </Button>
      </Form>
    </div>
  );
}

export default Home;
