import React, { useState,useEffect } from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import PersonIcon from '@material-ui/icons/Person';
import DeleteIcon from '@material-ui/icons/Delete';
import AddIcon from '@material-ui/icons/Add';
import Typography from '@material-ui/core/Typography';
import { useHistory } from 'react-router-dom';


export default function FriendRequests(){
    const [requestsArray,setRequestsArray] = useState([]);

    let history = useHistory();

    const removeRequest = (requestInfo) => {
        const data = {professionalId2:requestInfo.id}
        fetch('http://localhost:8080/v1/LinkedIn/removeFriendRequest', {
          method: "POST",
          mode:"cors",
          credentials:"include",
          body: JSON.stringify(data),
          headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
          })
          .then(response => response.json())
              .then((json) => {
                  if(json.error){
                      //Show error message
                      console.log(json.error);
                  }else{
                      console.log(json);
                      let newRequestsArray = requestsArray.filter((request) => request.id !== requestInfo.id);
                      setRequestsArray(newRequestsArray);
                  }
            });
    };

    const acceptRequest = (requestInfo) => {
        const data = {professionalId2:requestInfo.id}
        fetch('http://localhost:8080/v1/LinkedIn/addFriend', {
          method: "POST",
          mode:"cors",
          credentials:"include",
          body: JSON.stringify(data),
          headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
          })
          .then(response => response.json())
              .then((json) => {
                  if(json.error){
                      //Show error message
                      console.log(json.error);
                  }else{
                      console.log(json);
                      let newRequestsArray = requestsArray.filter((request) => request.id !== requestInfo.id);
                      setRequestsArray(newRequestsArray);
                  }
            });
    }
  
    useEffect(() => {
      fetch('http://localhost:8080/v1/LinkedIn/friendRequests', {
        method: "GET",
        mode:"cors",
        credentials:"include",
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
      })
      .then(response => response.json())
      .then((json) => {
        if(json.error){
          //Show error message
          console.log(json.error);
        }else{
          //Add the education info on the screen
          console.log(json.requests);
          if (json.requests!==null){
            setRequestsArray(json.requests);
          }
        }
      });    
    },[]);
  
    return(
        <React.Fragment>
        <Typography variant="h4" gutterBottom>
            Friend Requests
        </Typography>
        {requestsArray.length > 0 ?
        <List dense={false}>
            {requestsArray.map((request,index) =>{
                return(
                    <ListItem button onClick={() => {history.push(`professionals/${request.email}`)}} key={index}>
                    <ListItemAvatar>
                      <Avatar>
                        <PersonIcon />
                      </Avatar>
                    </ListItemAvatar>
                    <ListItemText
                      primary={request.firstName + " " + request.lastName}
                    />
                    <ListItemSecondaryAction>
                    <IconButton edge="end" aria-label="add" onClick={() => acceptRequest(request)}>
                        <AddIcon />
                      </IconButton>
                      <IconButton edge="end" aria-label="delete" onClick={() => removeRequest(request)}>
                        <DeleteIcon />
                      </IconButton>
                    </ListItemSecondaryAction>
                  </ListItem>
                );
                })}
        </List>
        :
        <Typography variant="h6" gutterBottom>
            No new friend requests
        </Typography>
        }
        </React.Fragment>
    )
}