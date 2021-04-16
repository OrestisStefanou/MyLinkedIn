import React, { useState } from 'react';
import ReactDom from 'react-dom';
import SignUp from "./SignUp"

function Greeting(){
    const [image, setImage] = useState({});
    const [userInfo,setUserInfo] = useState({email:'',first_name:'',last_name:'',password:'',password2:'',phone_number:''})

    const handleChange = (e) => {
        const name = e.target.name;
        const value = e.target.value;
        setUserInfo({ ...userInfo, [name]: value });
    };

    return <h4>Hello motherfuckers</h4>;
}

ReactDom.render(<SignUp/>,document.getElementById('root'));
