import React,{useState,useEffect} from 'react';
import JobAd from "./JobAd";

export default function JobAds() {
    const [ads,setAds] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/v1/LinkedIn/jobAds', {
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
                    console.log(json.jobAds);
                    if (json.jobAds!==null){
                        setAds(json.jobAds);
                    }
                }
            });    
    },[]);

    return(
        <React.Fragment>
            {ads && ads.map((ad) => {
                return(
                    <JobAd key={ad.id} adInfo={ad}/>
                )
            })}
        </React.Fragment>
    )
}