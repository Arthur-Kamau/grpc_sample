import { useNavigate } from "@solidjs/router";
import { Component, createSignal } from "solid-js"
import { appState } from "../../store/AppStore";
import toast from "solid-toast";

const DashboardPage: Component = () => {
    const navigate = useNavigate();
    const [userName, setUserName] = createSignal('');
    const [userNameResponse, setUserNameResponse] = createSignal('');
    return (
        <>
            <div class="view">
                <h1>Dashboard</h1>

                <form class="form">
                    <div class="mb-4">
                        <label class="form-label">Name Address</label>
                        <input
                            placeholder="someone@xyz.com" type="email" class="form-control"
                            value={userName()} onChange={(e) => setUserName(e.target.value)} />
                    </div>
                    <input type="submit" class="btn btn-primary form-control"
                        value="Submit"
                        onClick={async (e) => {
                            e.preventDefault();
                            console.log("hello wololo")
                            if (appState.client == undefined) {
                                alert("System not Available, please try again later.")
                            }
                            try {


                                const helloResponse = await appState.client?.sayHello({
                                    name: userName()
                                })

                                console.log("login response " + JSON.stringify(helloResponse?.response))
                                if (helloResponse?.response.isSuccess) {

                                    toast.error("Response =" + helloResponse?.response.sentence)
                                    setUserNameResponse( helloResponse?.response.sentence)
                                } else {
                                    toast.error(helloResponse?.response.reason)
                                    console.log("login response " + helloResponse)
                                }

                            } catch (e) {
                                toast.error("A critical error occurred (Connection Error)")
                                console.log("login error " + JSON.stringify(e))
                            }
                        }
                        } />
                </form>

                <div>Result : {userNameResponse()}</div>
            </div>
        </>)
}

export default DashboardPage;