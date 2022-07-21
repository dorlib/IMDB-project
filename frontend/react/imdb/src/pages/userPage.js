import React from 'react'
import {gql, useQuery} from "@apollo/client";
import Card from "../components/ui/Card";

import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import PeopleAltIcon from '@mui/icons-material/PeopleAlt';
import FiberManualRecordIcon from '@mui/icons-material/FiberManualRecord';
import Typography from "@mui/material/Typography";
import LocationOnIcon from '@mui/icons-material/LocationOn';
import SettingsSuggestIcon from '@mui/icons-material/SettingsSuggest';
import IconButton from '@mui/material/IconButton';
import EditIcon from '@mui/icons-material/Edit';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';

import classes from "./userPage.module.css"
import styled from "styled-components";
import DirectorsByUser from "../components/directors/directorsByUser";

function UserPage(props) {
    const USER_DATA = gql`
        query UserByID($id : ID!) {
            userById(id: $id) {
                id
                firstname
                lastname
                nickname
                description
                email
                birthday
                profile
                country
            }
        }
    `;


    const Details = styled.div`
        color: white;
        position: absolute;
        display: flex;
        right: 27.3cm;
        margin-top: 0.65cm;
    `;

    let TextBox = styled.text` 
        position: absolute;
        display: none;
        margin-top: 1cm;
        right: 21.9cm;
        background: #fff;
        font-size: small;
        ${Details}:hover & {
            display: flex;
            left: 1cm;
            top: -0.5cm;
            width: 4.5cm;
            color: black;
        }    
    `;

    let url = JSON.stringify(window.location.href);
    let lastSegment = parseInt(url.split("/").pop(), 10);

    const {loading, error, data} = useQuery(USER_DATA,
        {
            variables: {
                id: lastSegment || 0
            }

        })

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error : {error}</p>;

    let id = lastSegment
    let firstName = data["userById"]["0"]["firstname"]
    let lastName = data["userById"]["0"]["lastname"]
    let nickName = data["userById"]["0"]["nickname"]
    let Description = data["userById"]["0"]["description"]
    let Email = data["userById"]["0"]["email"]
    let Birthday = data["userById"]["0"]["birthday"]
    let Country = data["userById"]["0"]["country"]
    let Profile = data["userById"]["0"]["profile"] || "https://hope.be/wp-content/uploads/2015/05/no-user-image.gif"

    const handleClick = () => {
        return (
            window.location.replace("/editUserDetails/" + JSON.stringify(id))
        )
    }

    let Birth = Birthday.slice(2, 4) + '.' + Birthday.slice(0, 2) + '.' + Birthday.slice(4, 8)
    const getAge = birthDate => Math.floor((new Date() - new Date(birthDate).getTime()) / 3.15576e+10)

    let loaded = (
        <Card>
            <p className={classes.name}>
                {firstName} {lastName}
            </p>
            <p className={classes.nick}>
                {nickName}
            </p>
            <img className={classes.image} src={Profile} alt={"none"}/>
            <Stack direction="row" spacing={1} className={classes.changeProfile}>
                <IconButton color="primary" aria-label="change image" style={{color: "yellow"}}>
                    {props.LoggedInUser === id ? <SettingsSuggestIcon style={{fontSize: "xx-large"}}/> : null}
                </IconButton>
            </Stack>
            <PeopleAltIcon className={classes.followIcon}/>
            <Typography className={classes.follows}>Following</Typography>
            <FiberManualRecordIcon className={classes.dot}/>
            <Typography className={classes.followers}>Followers</Typography>
            <LocationOnIcon className={classes.location}/>
            <Typography className={classes.country}>{Country}</Typography>
            <Stack direction="row" spacing={2} className={classes.edit}>
                {props.LoggedInUser === id ? <Button variant="edit" className={classes.editBut} onClick={handleClick}>Edit
                    Profile</Button> : null}
            </Stack>
            <React.Fragment>
                <CardContent style={{backgroundColor: "darkslategray"}} className={classes.mainCard}>
                    <IconButton color="primary" aria-label="update details" className={classes.editDetails}>
                        {props.LoggedInUser === id ? <EditIcon className={classes.editDetailsBut}/> : null}
                        <Details>
                            <TextBox>
                                Click To Update Your Details
                            </TextBox>
                        </Details>
                    </IconButton>
                    <Typography variant="h5" component="div">
                        Hi, i'm {firstName}
                    </Typography>
                    <span className={classes.desc}>
                        If I Had To Describe Myself I Would Say : {Description}
                    </span>
                    <span className={classes.fav}>
                        My Favorite Movies Are : ...
                    </span>
                    <span className={classes.birth}>
                        I Was Born In {Birth} So Im {getAge(Birth)} Years Old
                    </span>
                    <span className={classes.contact}>
                        Wants To Contact {firstName} ? {Email}
                    </span>
                </CardContent>
            </React.Fragment>
            <React.Fragment>
                <CardContent style={{backgroundColor: "darkslategray"}} className={classes.pined}>
                    <Typography variant="h5" component="div">
                        Pined Movies Of {firstName}
                    </Typography>
                </CardContent>
            </React.Fragment>
            <React.Fragment>
                <CardContent style={{backgroundColor: "darkslategray"}} className={classes.Achievements}>
                    <Typography variant="h5" component="div">
                        Achievements Of {firstName}
                    </Typography>
                </CardContent>
            </React.Fragment>
            <React.Fragment>
                <CardContent style={{backgroundColor: "darkslategray"}} className={classes.contribMov}>
                    <Typography variant="h5" component="div">
                        Movies Contributed By {firstName}
                    </Typography>
                </CardContent>
            </React.Fragment>
            <React.Fragment>
                <CardContent style={{backgroundColor: "darkslategray"}} className={classes.contribDirec}>
                    <Typography variant="h5" component="div">
                        Directors Contributed By {firstName}
                    </Typography>
                    <DirectorsByUser userID={props.LoggedInUser}/>
                </CardContent>
            </React.Fragment>
            <React.Fragment>
                <CardContent style={{backgroundColor: "darkslategray"}} className={classes.contribRev}>
                    <Typography variant="h5" component="div">
                        Top Reviews By {firstName}
                    </Typography>
                </CardContent>
            </React.Fragment>
        </Card>

    )

    return loaded
}

export default UserPage