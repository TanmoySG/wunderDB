import React from 'react';
import { Container, Grid, Link, Hidden, Typography, Button } from "@material-ui/core";
import { faShieldAlt } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGithub } from '@fortawesome/free-brands-svg-icons';
import GoogleFontLoader from 'react-google-font-loader';



function Navbar() {
    return (
        <Container>
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
            <Grid container justify="center" alignItem="center" style={{ padding: "10px" }}>
                <Grid item xs={10} sm={6} md={3} lg={3} xl={3}>
                    <img src={require("../assets/wdb-complete.png")} style={{ height: "100%", width: "100%" }} alt="logo" />
                </Grid>
            </Grid>
            <Grid container justify="center" alignItem="center" style={{ padding: "10px" }}>
                <Grid item xs={12} sm={12} md={10} lg={10} xl={10}>
                    <Typography variant="h5" style={{ fontFamily: "Work Sans", textAlign: "center" }}>wunderDB is a JSON-based micro Document Database, inspired by <Link underline="always" href="https://www.mongodb.com" target="_blank" style={{ color: "#00d263" }}>MongoDB</Link>.</Typography>
                </Grid>
            </Grid>
            <Grid container justify="center" alignItem="center" style={{ padding: "10px" }} spacing={3}>
                <Link variant="subtitle1" underline="none" href="#registration" style={{ fontFamily: "DM Sans", fontWeight: "400", backgroundColor: "#A6FFCB", color: "#307D51", padding: "10px", borderRadius: "5px", margin: "5px" }}>Get-started</Link>
                <Link variant="subtitle1" underline="none" href="https://github.com/TanmoySG/wunderDB/blob/master/documentation/documentation.md" target="_blank" style={{ fontFamily: "DM Sans", fontWeight: "400", backgroundColor: "#C3DFFF", color: "#395a7f", padding: "10px", borderRadius: "5px", margin: "5px" }}>Documentation</Link>
                <Link variant="subtitle1" underline="none" href="https://github.com/TanmoySG/wunderDB/blob/master/documentation/deep-dive.md" target="_blank" style={{ fontFamily: "DM Sans", fontWeight: "400", backgroundColor: "#e5c7ff", color: "#734999", padding: "10px", borderRadius: "5px", margin: "5px" }}>Deep Dive</Link>
            </Grid>
        </Container>
    );
}

export default Navbar;
