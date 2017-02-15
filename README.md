# Getting started with Go on Bluemix
To get started, we'll take you through a sample Go hello world app, help you set up a development environment and deploy to {{site.data.keyword.Bluemix}}.

## Prerequisites
{: #prereqs}

You'll need the following:
* [{{site.data.keyword.Bluemix_notm}} account](https://console.ng.bluemix.net/registration/)
* [Cloud Foundry CLI ![External link icon](../../icons/launch-glyph.svg "External link icon")](https://github.com/cloudfoundry/cli#downloads){: new_window}
* [Git ![External link icon](../../icons/launch-glyph.svg "External link icon")](https://git-scm.com/downloads){: new_window}
* [Go ![External link icon](../../icons/launch-glyph.svg "External link icon")](https://golang.org/dl/){: new_window}

## 1. Clone the sample app
{: #clone}

Now you're ready to start working with the simple Go *hello world* app. Clone the repository and change to the directory where the sample app is located.
  ```
git clone https://github.com/IBM-Bluemix/get-started-go
  ```
  {: pre}
  ```
cd get-started-go
  ```
  {: pre}

  Peruse the files in the *get-started-go* directory to familiarize yourself with the contents.

## 2. Run the app locally
{: #run_locally}

  Build and run the app.
  ```
go run get-started-go.go
  ```
  {: pre}

  View your app at: http://localhost:8080

## 3. Prepare the app for deployment
{: #prepare}

To deploy to {{site.data.keyword.Bluemix_notm}}, it can be helpful to set up a manifest.yml file. One is provided for you with the sample. Take a moment to look at it.

The manifest.yml includes basic information about your app, such as the name, how much memory to allocate for each instance and the route. In this manifest.yml **random-route: true** generates a random route for your app to prevent your route from colliding with others.  You can replace **random-route: true** with **host: myChosenHostName**, supplying a host name of your choice. [Learn more...](/docs/manageapps/depapps.html#appmanifest)
 ```
 applications:
 - name: GetStartedGo
   random-route: true
   memory: 128M
 ```
 {: codeblock}

## 4. Deploy the app
{: #deploy}
You can use the Cloud Foundry CLI to deploy apps.

Choose your API endpoint
   ```
cf api <API-endpoint>
   ```
   {: pre}

Replace the *API-endpoint* in the command with an API endpoint from the following list.

|URL                             |Region          |
|:-------------------------------|:---------------|
| https://api.ng.bluemix.net     | US South       |
| https://api.eu-gb.bluemix.net  | United Kingdom |
| https://api.au-syd.bluemix.net | Sydney         |

Login to your {{site.data.keyword.Bluemix_notm}} account

  ```
cf login
  ```
  {: pre}

From within the *get-started-go* directory push your app to {{site.data.keyword.Bluemix_notm}}
  ```
cf push
  ```
  {: pre}

This can take a minute. If there is an error in the deployment process you can use the command `cf logs <Your-App-Name> --recent` to troubleshoot.

When deployment completes you should see a message indicating that your app is running.  View your app at the URL listed in the output of the push command.  You can also issue the

 ```
cf apps
 ```
  {: pre}
  command to view your apps status and see the URL.

Remember if you don't need your app live, stop it so you don't incur any unexpected charges.
{: tip}
