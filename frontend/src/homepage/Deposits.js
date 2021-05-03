import React,{useEffect,useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Title from './Title';

function preventDefault(event) {
  event.preventDefault();
}

const useStyles = makeStyles({
  depositContext: {
    flex: 1,
  },
});

export default function Deposits() {
  const classes = useStyles();
  const [professionalInfo,setProfessionalInfo] = useState({});

  useEffect(() => {
    const checkSession = async () => {
      const response = await fetch('http://localhost:8080/v1/LinkedIn/authenticated',{
        method: "GET",
        mode:"cors",
        credentials:"include",
        headers: {"Content-type": "application/json; charset=UTF-8",/*"Origin":"http://localhost:3000"*/}
        });
      const jsonResponse = await response.json();
      console.log(jsonResponse);
      if (response.status !== 202) {
        console.log("Something went wrong")
      }else{
        console.log(jsonResponse.professional);
        setProfessionalInfo(jsonResponse.professional);
      }
    };
    checkSession();
  },[]);

  return (
    <React.Fragment>
      <Title>{professionalInfo.firstName + " " + professionalInfo.lastName}</Title>
      
      <Title> Email:{professionalInfo.email} </Title>
      <Title> Phone Number:{professionalInfo.phoneNumber}</Title>
    </React.Fragment>
  );
}