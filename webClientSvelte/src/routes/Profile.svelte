<script lang="ts">
    import { onMount } from "svelte";
    import { deleteAccount, getUserInfo, updateEmail, updateInfos, updatePassword } from "../api/auth";
    import { pageTitle, userInfos } from "../api/stores";

    let username: string = '';
    let id: number;
    let kind: string = '';
    let email: string = '';
    let first_name: string = '';
    let last_name: string = '';

    let newEmail: string = '';
    let password: string = '';
    let newPassword: string = '';
    let currentPassword: string = '';
    let newFirstName: string = '';
    let newLastName: string = '';
    let newUsername: string = '';

    let emailError: string = '';
    let passwordError: string = '';
    let profileError: string = '';

  userInfos.subscribe(value => {
    if (value) {
      username = value.Username.String;
      id = value.ID;
      kind = value.Kind;
      email = value.Email;
      first_name = value.FirstName.String;
      last_name = value.LastName.String;
    }
  });

  function updateUserEmail() {
    if (!newEmail || !password) {
      emailError = "Information incomplete";
      return;
    }
    updateEmail(newEmail, password);
  };

  function updateUserPassword() {
    if (!newPassword || !currentPassword) {
      passwordError = "Information incomplete";
      return;
    }
    updatePassword(currentPassword, newPassword);
  };

  function updateProfile() {
    if (!newFirstName && !newLastName && !newUsername) {
      profileError = "At least one field must be filled";
      return;
    }
    updateInfos(newFirstName || undefined, newLastName || undefined, newUsername || undefined);
    getUserInfo();
  };

  function deleteThisAccount() {
    deleteAccount();
  }

  onMount(() => {
    pageTitle.set("Informations");
  });
</script>

<div class="profile-page">
    <div class="profile-container">
        <h1 class="username">username:<br/>{username}</h1>
        <h1 class="firstName">first name:<br/>{first_name}</h1>
        <h1 class="lastName">last name:<br/>{last_name}</h1>
        <h1 class="email">email:<br/>{email}</h1>
        <h1 class="id">account id:<br/>{id}</h1>
        <h1 class="kind">account type:<br/>{kind}</h1>
    </div>
    <div class="edit-container">
        <div>
            <h1>Update email</h1>
            <div class="form-container">
                <div class="input-fields">
                    <input placeholder="New Email" bind:value={newEmail} />
                    <input placeholder="Password" type="password" bind:value={password} />
                </div>
                <button on:click={updateUserEmail}>Update Email</button>
            </div>
            {#if emailError}
                <p class="error">{emailError}</p>
            {/if}
        </div>
        <div class="separator" />
        <div class="form">
            <h1>Update password</h1>
            <div class="form-container">
                <div class="input-fields">
                    <input placeholder="New Password" type="password" bind:value={newPassword} />
                    <input placeholder="Current Password" type="password" bind:value={currentPassword} />
                </div>
                <button on:click={updateUserPassword}>Update Password</button>
            </div>
            {#if passwordError}
                <p class="error">{passwordError}</p>
            {/if}
        </div>
        <div class="separator" />
        <div>
            <h1>Update informations</h1>
            <div class="form-container">
                <div class="input-fields">
                    <input placeholder="First Name" bind:value={newFirstName} />
                    <input placeholder="Last Name" bind:value={newLastName} />
                    <input placeholder="Username" bind:value={newUsername} />
                </div>
                <button on:click={updateProfile}>Update Profile</button>
            </div>
            {#if profileError}
                <p class="error">{profileError}</p>
            {/if}
        </div>
        <div class="separator" />
        <div>
          <h1>Delete Account</h1>
          <button on:click={deleteThisAccount}>Delete</button>
        </div>
    </div>
</div>

<style>
  .profile-page {
    width: 100%;
    display: flex;
    justify-content: space-around;
    align-items: center;
    flex-direction: column;
    padding: 1rem;
  }

  .profile-container {
    width: 80%;
    height: 100%;
    color: var(--color6);
    background-color: var(--color1);
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    font-family: 'lekton';
    border-radius: var(--borderRadius);
  }

  .profile-container h1 {
    background-color: var(--color1);
    padding-left: 2rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .edit-container {
    display: flex;
    flex-direction: column;
    width: 40%;
    height: 100%;
    justify-content: space-around;
    align-items: center;
    width: 100%;
    margin-bottom: 2rem;
  }

  .edit-container div {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    gap: 1rem;
    justify-content: center;
    align-items: center;
  }

  .edit-container .separator {
    width: 100%;
    height: 2px;
    margin: 1rem 0;
    background-color: var(--color6);
  }

  .edit-container .form-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .edit-container .input-fields {
    display: flex;
    flex-direction: column;
    width: 100%;
  }

  .edit-container div h1 {
    margin: 0;
  }

  .edit-container div input {
    padding: 0.5rem;
    border: none;
    border-radius: var(--borderRadius);
    font-size: 1rem;
    color: var(--color6)
  }

  .edit-container div input:focus {
    outline: none;
  }

  .edit-container div input::placeholder {
    color: var(--color6);
  }

  .edit-container div button {
    width: 150px;
    height: 50px;
    border: none;
    border-radius: var(--borderRadius);
    cursor: pointer;
    background-color: var(--color1);
    color: var(--color6);
    font-size: 1rem;
  }

  .error {
    color: red;
  }

    @media (min-width: 768px) {
      .profile-page {
        flex-direction: row;
        padding: 2rem;
      }

      .profile-container, .edit-container {
        width: 45%;
        margin-bottom: 0;
      }

      .edit-container .form-container {
        flex-direction: row;
        justify-content: space-between;
      }

      .edit-container .input-fields {
        width: 45%; /* Adjusts the input fields width */
      }

      .edit-container div button {
        width: auto; /* Adjust button width */
      }
    }

    @media (min-width: 1024px) {
        .profile-container, .edit-container {
            width: 40%;
        }
    }
</style>
