import React from 'react';
import './App.css';
import { Client as Styletron } from 'styletron-engine-atomic';
import { Provider as StyletronProvider } from 'styletron-react';
import { LightTheme, BaseProvider } from 'baseui';
import GoogleFontLoader from 'react-google-font-loader';
import { Link, Typography, Grid } from "@material-ui/core";
import { faShieldAlt } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGithub } from '@fortawesome/free-brands-svg-icons';
import Navbar from "./components/navbar.component";
import Hero from "./components/hero.component";
import GetStarted from "./components/get-started.component";


const engine = new Styletron();


function App() {

  return (

    <StyletronProvider value={engine}>
      <BaseProvider theme={LightTheme}>
        <GoogleFontLoader
          fonts={[
            {
              font: 'DM Sans',
              weights: [400, 500, 700],
            },
            {
              font: 'Work Sans',
              weights: [100, 200, 300, 400, 500, 600, 700, 800, 900],
            },
          ]}
        />
        <Grid container>
          <Grid container direction="row" alignItems="center" style={{ width: "100vw", height: "12vh", boxShadow: "-1px 12px 70px -1px rgba(161,161,161,0.15)" }} >
            <Navbar />
          </Grid>
          <Grid container direction="row" alignItems="center" style={{ height: "88vh" }} >
            <Hero />
          </Grid>
        </Grid>
        <section id="registration">
          <Grid container direction="row" alignItems="center" style={{ height: "100vh", padding: "20px" }} >
            <GetStarted />
          </Grid>
        </section>
      </BaseProvider>
    </StyletronProvider>
  );
}

export default App;
