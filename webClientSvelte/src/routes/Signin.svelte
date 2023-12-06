<script lang="ts">
    import { onMount } from 'svelte';
    import { Link } from "svelte-routing";
    import GoogleLogo from 'virtual:icons/flat-color-icons/google';
    import DiscordLogo from 'virtual:icons/skill-icons/discord';
    import { updateInfos, userLogin, userSignup } from '../api/auth';

    const DISCORD_URL = `${import.meta.env.VITE_DISCORD_CLIENT_API}?client_id=${import.meta.env.VITE_DISCORD_CLIENT_ID}&scope=identify%20email&redirect_uri=${import.meta.env.VITE_OAUTH_REDIRECT_URI}&response_type=code&state=discord`;
    const GOOGLE_URL = `${import.meta.env.VITE_GOOGLE_CLIENT_API}?client_id=${import.meta.env.VITE_GOOGLE_CLIENT_ID}&access_type=offline&response_type=code&state=google&redirect_uri=${import.meta.env.VITE_OAUTH_REDIRECT_URI}&scope=https://www.googleapis.com/auth/userinfo.email`;
    // const GITHUB_URL = `${import.meta.env.VITE_GITHUB_CLIENT_API}?client_id=${import.meta.env.VITE_GITHUB_CLIENT_ID}&redirect_uri=${import.meta.env.VITE_OAUTH_REDIRECT_URI}&scope=user&state=github`;

    let isRightPanelActive = '';
    let username = '';
    let email = '';
    let password = '';
    let confPass = '';
    let message = '';
    let firstName = '';
    let lastName = '';

    onMount(async () => {
        if (window.location.hash === '#signup') {
            activateSignUp();
        }
    });

    function activateSignUp() {
        isRightPanelActive = 'right-panel-active';
    }

    function activateSignIn() {
        isRightPanelActive = '';
        message = '';
    }

    async function handleSignup() {
        message = await userSignup(email, password, confPass);
        if(message === "Signed up successfully") {
            await updateInfos(firstName, lastName, username);
            activateSignIn();
        }
    }

    async function handleLogin() {
        await userLogin(email, password, message);
    }
</script>

<div class="container" class:right-panel-active={isRightPanelActive}>
	<div class="form-container sign-up-container">
		<form on:submit|preventDefault={handleSignup}>
			<h1>Create Account</h1>
			<div class="social-container">
				<!-- <a href={GITHUB_URL} class="social"><GithubLogo style="font-size: 2rem"/></a> -->
				<a href={GOOGLE_URL} class="social"><GoogleLogo style="font-size: 2rem"/></a>
				<a href={DISCORD_URL} class="social"><DiscordLogo style="font-size: 2rem"/></a>
			</div>
			<span>or use your email for registration</span>
			<input bind:value={username} type="text" placeholder="Username" />
            <input bind:value={firstName} type="text" placeholder="First name" />
            <input bind:value={lastName} type="text" placeholder="Last name" />
            <input bind:value={email} type="email" placeholder="Email" />
            <input bind:value={password} type="password" placeholder="Password" />
            <input bind:value={confPass} type="password" placeholder="Confirm Password" />
            {#if message}
                <p class="message">{message}</p>
            {/if}
			<button class="submitbutton">Sign Up</button>
		</form>
	</div>
	<div class="form-container sign-in-container">
		<form on:submit|preventDefault={handleLogin}>
			<h1>Sign in</h1>
			<div class="social-container">
				<!-- <a href="#" class="social"><GithubLogo style="font-size: 2rem"/></a> -->
				<a href={GOOGLE_URL} class="social"><GoogleLogo style="font-size: 2rem"/></a>
				<a href={DISCORD_URL} class="social"><DiscordLogo style="font-size: 2rem"/></a>
			</div>
			<span>or use your account</span>
			<input bind:value={email} type="email" placeholder="Email" />
            <input bind:value={password} type="password" placeholder="Password" />
            {#if message}
                <p class="message">{message}</p>
            {/if}
			<a href="#">Forgot your password?</a>
			<button>Sign In</button>
		</form>
	</div>
	<div class="overlay-container">
		<div class="overlay">
            <div class="overlay-panel overlay-left">
                <h1>Hello, Friend!</h1>
				<p>Enter your personal details and start journey with us</p>
                <p>or</p>
				<button class="ghost" on:click={activateSignIn}>Sign In</button>
                <nav class="link-wrapper">
                    <Link class="link" to="/">&lt;--Back to Home</Link>
                </nav>
			</div>
            <div class="overlay-panel overlay-right">
                <h1>Welcome Back!</h1>
                <p>To keep connected with us please login with your personal info</p>
                <p>or</p>
                <button class="ghost" on:click={activateSignUp}>Sign Up</button>
                <nav class="link-wrapper">
                    <Link class="link" to="/">&lt;--Back to Home</Link>
                </nav>
            </div>
		</div>
	</div>
</div>

<style>
    * {
	    box-sizing: border-box;
    }

    h1 {
        font-weight: bold;
        margin: 0;
        color: var(--color2);
    }

    p {
        font-size: 20px;
        font-weight: 100;
        line-height: 20px;
        letter-spacing: 0.5px;
        margin: 20px 0 30px;
    }

    span {
        font-size: 18px;
        color: var(--color2);
    }

    a {
        color: #333;
        font-size: 14px;
        text-decoration: none;
        margin: 15px 0;
    }

    button {
        border-radius: var(--borderRadius);
        border: none;
        background: linear-gradient(to right, var(--color2), var(--color6));
        color: #FFFFFF;
        font-size: 20px;
        font-weight: bold;
        height: 3rem;
        width: 10rem;
        letter-spacing: 1px;
        text-transform: uppercase;
        transition: transform 80ms ease-in;
    }

    button:active {
        transform: scale(0.95);
    }

    button:focus {
        outline: none;
    }

    button.ghost {
        background: linear-gradient(to right, var(--color6), var(--backgroundColor));
        box-shadow: rgba(60, 64, 67, 0.3) 0px 1px 2px 0px, rgba(60, 64, 67, 0.15) 0px 2px 6px 2px;
    }

    form {
        background-color: var(--color1);
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        padding: 0 50px;
        height: 100%;
        text-align: center;
    }

    input {
        background-color: var(--color1);
        border: 1px solid black;
        padding: 0 15px;
        box-shadow: rgba(60, 64, 67, 0.3) 0px 1px 2px 0px, rgba(60, 64, 67, 0.15) 0px 2px 6px 2px;
        margin: 0.5rem 0;
        height: 3rem;
        width: 70%;
        border-radius: var(--borderRadius);
    }

    input::placeholder {
        color: var(--color2);
        font-weight: bold;
        font-size: 1.5rem;
        vertical-align: middle;
    }

    input::-webkit-input-placeholder {
        color: var(--color2);
        font-weight: bold;
        font-size: 1.5rem;
        vertical-align: middle;
    }

    input::-moz-placeholder {
        color: var(--color2);
        font-weight: bold;
        font-size: 1.5rem;
        vertical-align: middle;
    }

    .container {
        background-color: var(--color1);
        border-radius: var(--borderRadius);
        box-shadow: 0 14px 28px rgba(0,0,0,0.25), 0 10px 10px rgba(0,0,0,0.22);
        position: relative;
        overflow: hidden;
        width: 100%;
        height: 100%;
    }

    .form-container {
        position: absolute;
        top: 0;
        height: 100%;
        transition: all 0.6s ease-in-out;
    }

    .sign-in-container {
        left: 0;
        width: 50%;
        z-index: 2;
    }

    .container.right-panel-active .sign-in-container {
        transform: translateX(100%);
    }

    .sign-up-container {
        left: 0;
        width: 50%;
        opacity: 0;
        z-index: 1;
    }

    .submitbutton {
        margin-top: 2rem;
    }

    .container.right-panel-active .sign-up-container {
        transform: translateX(100%);
        opacity: 1;
        z-index: 5;
        animation: show 0.6s;
    }

    @keyframes show {
        0%, 49.99% {
            opacity: 0;
            z-index: 1;
        }

        50%, 100% {
            opacity: 1;
            z-index: 5;
        }
    }

    .overlay-container {
        position: absolute;
        top: 0;
        left: 50%;
        width: 50%;
        height: 100%;
        overflow: hidden;
        transition: transform 0.6s ease-in-out;
        z-index: 100;
    }

    .container.right-panel-active .overlay-container{
        transform: translateX(-100%);
    }

    .overlay {
        background: var(--color6);
        background: -webkit-linear-gradient(to right, var(--color6), #FF416C);
        background: linear-gradient(to right, var(--color6), #FF416C);
        background-repeat: no-repeat;
        background-size: cover;
        background-position: 0 0;
        color: #FFFFFF;
        position: relative;
        left: -100%;
        height: 100%;
        width: 200%;
        transform: translateX(0);
        transition: transform 0.6s ease-in-out;
    }

    .container.right-panel-active .overlay {
        transform: translateX(50%);
    }

    .overlay-panel {
        position: absolute;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        padding: 0 40px;
        text-align: center;
        top: 0;
        height: 100%;
        width: 50%;
        transform: translateX(0);
        transition: transform 0.6s ease-in-out;
    }

    .overlay-left {
        transform: translateX(-20%);
    }

    .container.right-panel-active .overlay-left {
        transform: translateX(0);
    }

    .overlay-right {
        right: 0;
        transform: translateX(0);
    }

    .container.right-panel-active .overlay-right {
        transform: translateX(20%);
    }

    .social-container {
        margin: 20px 0;
    }

    .social-container a {
        border: 1px solid #8d8d8d;
        border-radius: 50%;
        display: inline-flex;
        justify-content: center;
        align-items: center;
        margin: 0 5px;
        height: 60px;
        width: 60px;
    }

    .link-wrapper {
        margin-top: 5rem;
    }

    .link-wrapper > :global(a) {
        text-decoration: none;
        color: var(--backgroundColor)
    }
</style>
