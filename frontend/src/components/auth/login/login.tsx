/* Login */

import axios from "axios";
import { FormEventHandler, ChangeEventHandler, useEffect, useState } from "react";
import Cookies from "universal-cookie";

export default function Login() {
	/* Login Page */
	const url = new URL(window.location.href);
	const msg_success = url.searchParams.get("msg_success");

	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [error_status, setErrorStatus] = useState(false);
	const [success_status, setSuccessStatus] = useState(false);
	const [msg_error, setMsgError] = useState("");

	useEffect(() => {
		if (msg_success !== null) {
			setSuccessStatus(true);
		}
		if (error_status === true) {
			setSuccessStatus(false);
		}
	}, [msg_success, error_status]);

	const handleUsernameChange: ChangeEventHandler<HTMLInputElement> = e => {
		const target = e.target as HTMLInputElement;
		setUsername(target.value);
	};

	const handlePasswordChange: ChangeEventHandler<HTMLInputElement> = e => {
		const target = e.target as HTMLInputElement;
		setPassword(target.value);
	};

	const handleSubmit: FormEventHandler<HTMLFormElement> = async e => {
		e.preventDefault();
		axios.post("http://localhost:8080/api/login", {
			username: username,
			password: password,
		}).then((response) => {
			const cookies = new Cookies();
			if (response.data.auth === true) {
				// set cookie
				cookies.set("token", response.data.token);
                window.location.href = "/dashboard";
			}
			else {
				// display error message
				setErrorStatus(true);
				setMsgError(response.data.msg_error);
			}
		}).catch(e => {
            console.log(e);
        })
	};

	return (		
		<div>
			{ error_status === true &&
			<div className="alert alert-danger" role="alert">
				{ msg_error } 
			</div>
			}
			{ success_status === true &&
			<div className="alert alert-success" role="alert">
				{ msg_success } 
			</div>
			}
			<h2> Login </h2>
			<form method="post" onSubmit={handleSubmit}>
				<div className="form-group">
					<label htmlFor="exampleInputUsername1">Email or Username</label>
					<input className="form-control" id="exampleInputUsername1" 
						name="username" placeholder="Enter email or username" 
						value={username} 
						onChange={handleUsernameChange}
					required />
				</div>
				<div className="form-group">
					<label htmlFor="exampleInputPassword1">Password</label>
					<input type="password" className="form-control" id="exampleInputPassword1" 
						name="password" placeholder="Enter password" 
						value={password}
						onChange={handlePasswordChange}
					required />
				</div>
				<button type="submit" className="btn btn-primary">Submit</button>
			</form>
		</div>
	);
}
