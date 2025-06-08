
import { Component } from 'solid-js';

const Navbar: Component = () => {
  return (
    <div>
        <h2>login</h2>
        <fieldset class="fieldset">
  <legend class="fieldset-legend">email</legend>
  <input type="text" class="input" placeholder="Type here" />
  <p class="label">Optional</p>
</fieldset>
<fieldset class="fieldset">
  <legend class="fieldset-legend">password</legend>
  <input type="text" class="input" placeholder="Type here" />
  <p class="label">Optional</p>
</fieldset>
<button class="btn">Default</button>
    </div>
  );
};

export default Navbar;