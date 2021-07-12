import { useEffect, useState } from 'react';
import { LinearProgress, Container } from '@material-ui/core';
import { Alert } from '@material-ui/lab';
import axios from 'axios';
import { config } from './config';

const styles = {
  alertSection: {
    margin: "1rem"
  }
}

function Status() {
  const [data, setData] = useState({});
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);
  useEffect(() =>
    axios.get(`${config.baseUrl}/v1/health`)
      .then(({ data }) => setData(data))
      .catch(({ message }) => setError(message))
      .finally(() => setLoading(false)), [])

  if (loading) return <LinearProgress />
  if (error !== "") return <Alert severity="error">{error}</Alert>
  return (<>
    <div style={styles.alertSection}>
      <Alert severity={data.message === "up" ? "success" : "error"}>Overall Network</Alert>
    </div>
    {Object.keys(data.services).map(service => {
      const status = data.services[service];
      return (<div style={styles.alertSection}>
        <Alert severity={status === "up" ? "success" : "error"}> {service} </Alert>
      </div>);
    })}
  </>);
}

function App() {
  return (<>
    <Container>
      <Status />
    </Container>
  </>);
}

export default App;
