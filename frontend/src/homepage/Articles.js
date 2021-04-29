import React,{useState,useEffect} from 'react';
import Article from "./Article";

export default function Articles() {
    const [articles,setArticles] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/v1/LinkedIn/getArticles', {
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
                    console.log(json.articles);
                    if (json.articles!==null){
                    setArticles(json.articles);
                    }
                }
            });    
    },[]);

    return(
        <React.Fragment>
            {articles && articles.map((article) => {
                return(
                    <Article key={article.id} articleInfo={article}/>
                )
            })}
        </React.Fragment>
    )
}