<div class='showcase-container'>
    <div class="container">
        <h1 class="pricing-title">Plans</h1>
        <div class="container-box">
            <div class="box">
                <h1>Starter</h1>
                <div class="price-box">
                    <p class="start-price">Starts at</p>
                    <h1 class="price">$10 / month</h1>
                </div>
                <div class="centered-list">
                    <ul>
                        <li>5 services</li>
                        <li>Exclusive triggers and actions</li>
                        <li>Customer support</li>
                    </ul>
                </div>
                <div class="button-container">
                    <button class="bouton">Subscribe Now</button>
                </div>
            </div>
            <div class="box">
                <h1>Professional</h1>
                <div class="price-box">
                    <p class="start-price">Starts at</p>
                    <h1 class="price">$20 / month</h1>
                </div>
                <div class="centered-list">
                    <ul>
                        <li>All services</li>
                        <li>Connect multiple accounts</li>
                        <li>AI Services</li>
                        <li>Developer tools</li>
                        <li>Prioritized customer support</li>
                    </ul>
                </div>
                <div class="button-container">
                    <button class="bouton">Subscribe Now</button>
                </div>
            </div>
        </div>

    </div>
</div>

<style>
    .showcase-container {
        position: relative;
        width: 100%;
        height: 100%;
        background-color: var(--backgroundColor);
    }

    h1 {
        margin: 0;
        color: var(--color2);
    }

    .pricing-title {
        font-weight: bold;
        font-size: 54px;
        margin: 0;
        color: var(--color2);
    }

    .container {
        position: relative;
        text-align: center;
        height: 100vh;
        overflow: hidden;
    }

    .pricing-title {
        margin-top: 10rem;
    }

    .container-box {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        max-width: 1024px;
        margin: 1rem auto;
    }

    .box {
        flex: 1 1 calc(45% - 20px);
        margin: 0 10px 20px;
        background-color: var(--color1);
        padding: 20px;
        border: 1px solid var(--color2);
        border-radius: 20px;
        display: flex;
        flex-direction: column;
    }

    .price-box {
        margin-top: 0.5rem;
        flex-grow: 1;
    }

    p {
        font-weight: bold;
    }

    .start-price {
        margin: 0;
        font-size: 20px;
    }


    button {
        border-radius: var(--borderRadius);
        border: none;
        background: linear-gradient(to right, var(--color2), var(--color6));
        color: #FFFFFF;
        font-size: 20px;
        font-weight: bold;
        height: 3rem;
        width: 20rem;
        letter-spacing: 1px;
        text-transform: uppercase;
        transition: transform 0.4s ease-in-out;
    }

    button:active {
        transform: scale(0.95);
    }

    button:hover {
        cursor: pointer;
        transform: scale(1.10);
        transition: transform 0.35s ease-in-out;
    }

    .price {
        margin: 0;
        font-size: 40px;
    }

    .centered-list {
        text-align: center;
        flex-grow: 1;
        padding-top: 1.5rem;
    }

    .centered-list ul {
        list-style-type: none;
        padding: 0;
        font-size: 22px;
        margin: 0;
        flex-grow: 1;
    }

    .centered-list li {
        display: block;
        margin: 0 10px;
    }

    .button-container {
        display: flex;
        justify-content: center;
        margin-top: 2rem;
    }

    @media (max-width: 768px) {
        .container-box {
            max-width: 424px;
        }

        .box {
            flex: 0 0 calc(100% - 20px);
            margin: 0 0 20px;
        }
    }

</style>
