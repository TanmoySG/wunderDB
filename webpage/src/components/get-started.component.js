import React, { useState } from 'react';
import { useStyletron } from 'baseui';
import { Container, Grid, Paper, Typography, Tooltip, Link, Badge } from "@material-ui/core";
import { faCubes, faCopy, faUser, faLock, faAt, faLayerGroup } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { FormControl } from "baseui/form-control";
import { Input } from "baseui/input";
import { Alert } from 'baseui/icon';
import { CopyToClipboard } from 'react-copy-to-clipboard';
import { validate as validateEmail } from 'email-validator';
import { Button } from "baseui/button";
import { Textarea } from "baseui/textarea";
import 'status-indicator/styles.css'


function Negative() {
    const [css, theme] = useStyletron();
    return (
        <div
            className={css({
                display: 'flex',
                alignItems: 'center',
                paddingRight: theme.sizing.scale500,
                color: theme.colors.negative400,
            })}
        >
            <Alert size="18px" />
        </div>
    );
}

function GetStarted() {

    const [clusterId, setClusterId] = useState('<cluster-id>');
    const [tokens, setTokens] = useState(['<one-of-the-three-tokens-generated>']);
    const [requestResponse, setrequestResponse] = useState('Send Request');
    const [status, setStatus] = useState('2')
    const [jsonResponse, setjsonResponse] = useState({});
    const [name, setName] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [isValid, setIsValid] = React.useState(false);
    const [isVisited, setIsVisited] = React.useState(false);
    const shouldShowError = !isValid && isVisited;
    const onChange = ({ target: { value } }) => {
        setIsValid(validateEmail(value));
        setUsername(value);
    };
    var payload = {
        "name": name,
        "username": username,
        "password": password
    };
    var payloadPrint = "{ \n\t 'name' : '" + name + "', \n\t 'username' : '" + username + "', \n\t 'password' : '" + password + "' \n }";
    var uaaEndpoint = "https://wdb.tanmoysg.com/connect?cluster=" + clusterId + "&token=" + tokens[Math.floor(Math.random() * tokens.length)];

    const handleSubmit = (e) => {
        console.log(payload);
        e.preventDefault();
        fetch("https://wdb.tanmoysg.com/register", {
            method: "POST",
            cache: "no-cache",
            headers: {
                "content_type": "application/json",
            },
            body: JSON.stringify(payload)
        }).then(response => {
            return response.json()
        }).then(json => {
            setjsonResponse(json);
            console.log(json);
            setrequestResponse(json.response);
            if (json.status_code === '1') {
                setStatus(json.status_code);
                setClusterId(json.cluster_id);
                setTokens(json.access_tokens);
            } else {
                setStatus(json.status_code);
            }
        })
    };

    return (
        <Container >
            <Grid container direction="row" justify="center" alignItems="center">
                <Typography variant="h5" style={{ fontFamily: "DM Sans", fontWeight: "500", color: "#307D51", padding: "10px" }}><FontAwesomeIcon icon={faCubes} style={{ color: "#307D51", marginRight: "5px", marginLeft: "5px" }} /> Get Started </Typography>
                <Grid item xs={12} sm={12} md={12} lg={12} xl={12}>
                    <Grid container direction="row" justify="center" alignItems="center">
                        <Grid item xs={12} sm={12} md={6} lg={6} xl={6} justify="center" alignItems="center">
                            <Typography variant="subtitle1" style={{ fontFamily: "DM Sans", fontWeight: "500", color: "black", textAlign: "center" }}>Unified Actions API Endpoint</Typography>
                            <Input endEnhancer={
                                <CopyToClipboard text={uaaEndpoint}>
                                    <Tooltip title="Click to copy">
                                        <FontAwesomeIcon icon={faCopy} />
                                    </Tooltip>
                                </CopyToClipboard>
                            } value={uaaEndpoint} style={{ boxShadow: "0px 0px 32px -6px rgba(0,0,0,0.72)" }} />
                        </Grid>
                    </Grid>
                </Grid>
                <Grid container direction="row" spacing={3}>
                    <Grid item xs={12} sm={12} md={4} lg={4} xl={4} style={{ marginTop: "15px" }}>
                        <Paper style={{ padding: "20px" }} elevation={5}>
                            <form onSubmit={handleSubmit}>
                                <Typography variant="h5" style={{ fontFamily: "DM Sans", fontWeight: "500", color: "#395a7f" }}>Register</Typography>
                                <FormControl label="Name" >
                                    <Input startEnhancer={<FontAwesomeIcon icon={faUser} />} onChange={e => setName(e.target.value)} />
                                </FormControl>
                                <FormControl label="Email" error={shouldShowError ? 'Please input a valid email address' : null}>
                                    <Input startEnhancer={<FontAwesomeIcon icon={faAt} />} onChange={onChange} onBlur={() => setIsVisited(true)} error={shouldShowError} overrides={shouldShowError ? { After: Negative } : {}} type="email" required />
                                </FormControl>
                                <FormControl label="Password" >
                                    <Input startEnhancer={<FontAwesomeIcon icon={faLock} />} type="password" onChange={e => setPassword(e.target.value)} />
                                </FormControl>
                                <FormControl label="Payload" caption="Payload to be sent to the API." >
                                    <Textarea
                                        value={payloadPrint}
                                        disabled
                                    />
                                </FormControl>
                                <Button preventDefault onClick={handleSubmit} style={{ fontFamily: "DM Sans", fontWeight: "400", backgroundColor: "#C3DFFF", color: "#395a7f", padding: "10px", borderRadius: "5px", margin: "5px" }}> <FontAwesomeIcon icon={faLayerGroup} style={{ color: "#395a7f", marginRight: "5px", marginLeft: "5px" }} /> Create Cluster</Button>
                            </form>
                        </Paper>
                    </Grid>
                    <Grid item xs={12} sm={12} md={8} lg={8} xl={8} style={{ marginTop: "15px" }}>
                        <Paper style={{ padding: "20px" }} elevation={5}>
                            <Grid container direction="row" alignItems="center" justify="space-between">
                                <Grid item>
                                    <Typography variant="h5" style={{ fontFamily: "DM Sans", fontWeight: "500", color: "#395a7f"}}>Response</Typography>
                                </Grid>
                                <Grid item>
                                    <Tooltip title={status === '2' ? "No Request in Pipeline" : status === '1' ? "Request Successful" : "Request Unsuccessful"}>
                                        {status === '2' ? <status-indicator pulse></status-indicator> : status === '1' ? <status-indicator positive pulse></status-indicator> : <status-indicator negative pulse></status-indicator>}
                                    </Tooltip>
                                </Grid>
                            </Grid>
                            <Textarea
                                value={JSON.stringify(jsonResponse, undefined, 4)}
                                overrides={{
                                    Input: {
                                        style: {
                                            maxHeight: '300px',
                                            minHeight: '200px',
                                            minWidth: '300px',
                                            width: '100vw', // fill all available space up to parent max-width
                                        },
                                    },
                                    InputContainer: {
                                        style: {
                                            maxWidth: '100%',
                                            width: 'min-content',
                                        },
                                    },
                                }}
                                style={{overflow: "auto"}}
                            />
                            <Typography variant="subtitle1" style={{ fontFamily: "DM Sans", fontWeight: "400" }}> Response:</Typography>
                            <Typography variant="h6" style={{ fontFamily: "DM Sans", fontWeight: "600", color: "#395a7f", overflow: "auto" }}>{requestResponse}</Typography>
                            <Typography variant="subtitle1" style={{ fontFamily: "DM Sans", fontWeight: "400" }}> Cluster ID:</Typography>
                            <CopyToClipboard text={clusterId}>
                                <Tooltip title="Click to copy">
                                    <Typography variant="h6" style={{ fontFamily: "DM Sans", fontWeight: "600", color: "#395a7f" , overflow: "auto"}}>{clusterId}</Typography>
                                </Tooltip>
                            </CopyToClipboard>
                            <Typography variant="subtitle1" style={{ fontFamily: "DM Sans", fontWeight: "400" }}>Access Tokens:</Typography>

                            {
                                tokens.map((token) => {
                                    return (
                                        <CopyToClipboard text={token}>
                                            <Tooltip title="Click to copy">
                                                <Typography variant="h6" style={{ fontFamily: "DM Sans", fontWeight: "600", color: "#395a7f" ,overflow: "auto" }}>{token}</Typography>
                                            </Tooltip>
                                        </CopyToClipboard>
                                    )
                                })
                            }
                        </Paper>
                        <Grid container direction="row" justify="flex-end" alignItems="center" style={{ textAlign: "center", paddingTop : "20px" }} >
                            <Typography variant="subtitle1">Developed by <Link variant="subtitle1" underline="none" target="_blank" href="https://www.tanmoysg.com" style={{ fontFamily: "DM Sans", fontWeight: "400", color: "#734999" }}>Tanmoy Sen Gupta</Link></Typography>
                        </Grid>
                    </Grid>
                </Grid>
            </Grid>

        </Container>
    );
}

export default GetStarted;

