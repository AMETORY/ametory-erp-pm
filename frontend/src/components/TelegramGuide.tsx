
const TelegramIntegrationGuide: React.FC = () => {
    return (
      <div className="p-6 bg-gray-100  flex items-center justify-center my-6">
        <div className="bg-white rounded-lg shadow-lg max-w-2xl w-full p-8">
          <h2 className="text-2xl font-semibold text-gray-800 mb-4">Telegram Integration Guide</h2>
          <p className="text-gray-600 mb-6">
            Follow these steps to integrate your Telegram account with CRM. This will allow you to manage customer interactions directly through the platform.
          </p>
  
          <ol className="list-decimal list-inside space-y-4">
            <li>
              <strong>Create a Telegram Bot:</strong> Open the Telegram app and search for <span className="font-mono">@BotFather</span>. Send the command <span className="font-mono">/newbot</span> to create a new bot, and follow the prompts to name and set up your bot. Save the bot token provided.
            </li>
            <li>
              <strong>Get Your Bot Token:</strong> The token provided by BotFather is required for connecting the bot. Save it securely, as you'll need it for the next steps.
            </li>
            <li>
              <strong>Copy the Bot Token:</strong> Copy the bot token provided by BotFather securely, as you'll need it for the next steps.
            </li>
            <li>
              <strong>Paste Bot Token:</strong> at BOT TOKEN field.
            </li>
            <li>
              <strong>Save Connection:</strong> After Paste Bot Token, save the Connection settings.
            </li>
            <li>
              <strong>Test the Integration:</strong> Send a test message to your bot in Telegram to confirm that messages are routed correctly to the CRM platform.
            </li>
          </ol>
  
         
        </div>
      </div>
    );
  };
  
  export default TelegramIntegrationGuide;