# proxy6-net-tools

Utilities for working with proxy6.net api

# Getting started
First, create a file with .env, in it create a variable API_TOKEN, and give it a value (specify your api token)

If you don't have your own api token at the moment, you can release it here - https://proxy6.net/en/user/developers

# Working with API
The api is accessed via http requests to the proxy6.net server. Using these utilities, you can view your list of active proxies, find out your account ID, current account balance, check the number of available proxies by specifying a certain country

I have not implemented the functionality to buy proxies, as for me personally, this is not safe, and from some point of view is even pointless, because you still have to go to the control panel to replenish the balance (api does not provide the ability to replenish balance)

# FAQ
I tried to use all the features of Go to achieve the best possible code. However, if you manage to find a bug, or if you want to suggest any improvements, I am always ready to consider your suggestions in the

All methods and variables in the code are commented for developers, I hope there are no questions about them
