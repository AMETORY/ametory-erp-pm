import type { FC } from "react";

interface TermUsePageProps {}

const TermUsePage: FC<TermUsePageProps> = ({}) => {
  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gradient-to-r from-blue-500 to-purple-500 py-8">
      <div className="w-full max-w-3xl p-4 bg-white rounded-lg shadow-md overflow-y-auto">
        <div className="flex items-center justify-center mb-4 ">
          <div className="flex flex-col justify-center items-center">
            <img
              src="https://app.senandika.web.id/assets/static/logo-senandika.jpg"
              alt="Logo"
              className="w-24 h-24"
            />
            <div className="container mx-auto px-4 py-8 text-gray-700">
              <h1 className="text-3xl font-bold mb-6 text-center">
                Terms of Use
              </h1>
              <p>
                <em>Last updated: May 17, 2025</em>
              </p>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  1. Acceptance of Terms
                </h2>
                <p>
                  By accessing or using the <strong>Senandika</strong> service,
                  you agree that you have read, understood, and accepted all of
                  these terms. If you do not agree, please do not use this
                  service.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">2. Our Services</h2>
                <p>
                  Senandika provides a digital platform to manage Instagram
                  messages, automate DMs, and perform social media analytics. We
                  reserve the right to add, modify, or discontinue services at
                  any time without prior notice.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  3. User Accounts
                </h2>
                <p>
                  To access certain features, you may need to create an account.
                  You are responsible for maintaining the security of your login
                  information. We are not responsible for any loss resulting
                  from unauthorized access to your account.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">4. Privacy</h2>
                <p>
                  Your use of our services is also subject to our{" "}
                  <a href="/privacy" target="_blank" rel="noopener noreferrer">
                    Privacy Policy
                  </a>
                  , which explains how your data is used and protected.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  5. Intellectual Property Rights
                </h2>
                <p>
                  All content, trademarks, logos, and copyrights on this
                  platform are owned by Senandika or third parties who license
                  them to us. You may not copy or distribute content without
                  permission.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  6. User Obligations
                </h2>
                <ul>
                  <li>
                    Use the service lawfully and without violating any laws
                  </li>
                  <li>Do not send spam, misleading, or harmful content</li>
                  <li>Do not disrupt our systems, APIs, or services</li>
                </ul>
                <p>
                  We reserve the right to suspend or delete accounts that
                  violate these terms.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  7. Limitation of Liability
                </h2>
                <p>
                  Our service is provided “as is” without warranties. We are not
                  liable for direct or indirect losses arising from your use of
                  our services.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  8. Changes to Terms
                </h2>
                <p>
                  We may update these terms at any time. You will be notified
                  via the site or email. Continued use of the service means you
                  agree to those changes.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  9. Governing Law
                </h2>
                <p>These terms are governed by the laws of Indonesia.</p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">10. Contact Us</h2>
                <p>If you have any questions, please contact us:</p>
                <ul>
                  <li>
                    Email:{" "}
                    <a href="mailto:support@ametory.id">support@ametory.id</a>
                  </li>
                  <li>
                    Website:{" "}
                    <a
                      href="https://senandika.web.id"
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      senandika.web.id
                    </a>
                  </li>
                </ul>
              </section>
            </div>
          </div>
        </div>
      </div>
      <footer className="text-center text-sm mt-4 text-white">
        © 2024 Senandika. All rights reserved.
      </footer>
    </div>
  );
};
export default TermUsePage;
