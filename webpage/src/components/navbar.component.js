import React from 'react';
import { Container, Grid, Link, Hidden, Typography } from "@material-ui/core";
import { faShieldAlt } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGithub } from '@fortawesome/free-brands-svg-icons';


function Navbar() {
    return (
        <Container>
            <Grid container>
                <Grid item xs={12} sm={12} md={12} lg={12} xl={12}>
                    <img src={require("../assets/wdb.png")} style={{ height: "30px" }} alt="logo" />
                    <Link variant="h5" underline="none" href="https://github.com/TanmoySG/wunderDB"  target="_blank" style={{  float: "right", color: "#26242F" }}><FontAwesomeIcon icon={faGithub} style={{ color: "#26242F" }} /></Link>
                </Grid>
            </Grid>
        </Container>
    );
}

export default Navbar;
