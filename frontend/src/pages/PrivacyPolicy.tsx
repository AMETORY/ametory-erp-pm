import type { FC } from "react";

interface PrivacyPolicyProps {}

const PrivacyPolicy: FC<PrivacyPolicyProps> = ({}) => {
  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gradient-to-r from-blue-500 to-purple-500 py-8">
      <div className="w-full max-w-3xl p-4 bg-white rounded-lg shadow-md overflow-y-auto">
        <div className="flex items-center justify-center mb-4 ">
          <div className="flex flex-col justify-center items-center">
            <img src="https://app.senandika.web.id/assets/static/logo-senandika.jpg" alt="Logo" className="w-24 h-24" />
            <div className="container mx-auto px-4 py-8 text-gray-700">
              <h1 className="text-3xl font-bold mb-6 text-center">
                Privacy Policy of Senandika
              </h1>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">1. Introduction</h2>
                <p>
                  Senandika values your privacy and is committed to protecting
                  your personal data. This policy explains how we collect, use,
                  and share information when you use our platform.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  2. Information We Collect
                </h2>
                <p>We collect various types of information, including:</p>
                <ul className="list-disc list-inside ml-6">
                  <li>
                    <strong>Personal Information:</strong> Name, email, company
                    name, and other identifiers.
                  </li>
                  <li>
                    <strong>Usage Data:</strong> Information on how you interact
                    with Senandika.
                  </li>
                  <li>
                    <strong>Technical Data:</strong> IP address, browser type,
                    device information.
                  </li>
                </ul>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  3. How We Use Your Information
                </h2>
                <p>
                  Senandika uses your information for the following purposes:
                </p>
                <ul className="list-disc list-inside ml-6">
                  <li>Providing and improving our services.</li>
                  <li>Personalizing your experience on our platform.</li>
                  <li>
                    Communicating updates, promotions, and service information.
                  </li>
                  <li>For security and compliance with legal obligations.</li>
                </ul>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  4. Data Sharing and Disclosure
                </h2>
                <p>We only share your data under the following conditions:</p>
                <ul className="list-disc list-inside ml-6">
                  <li>
                    <strong>With Your Consent:</strong> When you explicitly
                    authorize sharing.
                  </li>
                  <li>
                    <strong>Service Providers:</strong> Third parties who help
                    us operate our platform.
                  </li>
                  <li>
                    <strong>Legal Requirements:</strong> When required by law or
                    for legal protection.
                  </li>
                </ul>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  5. Security of Your Information
                </h2>
                <p>
                  We employ industry-standard security measures to safeguard
                  your data. However, no system is completely secure, and we
                  cannot guarantee absolute security.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">6. Your Rights</h2>
                <p>
                  Depending on your location, you may have the following rights
                  regarding your data:
                </p>
                <ul className="list-disc list-inside ml-6">
                  <li>
                    <strong>Delete Account:</strong> You can delete your account
                    and all associated data by contacting our support team at
                    <a href="mailto:support@ametory.id" className="text-blue-600">
                      support@ametory.id
                    </a>
                    . This will erase all personal data and account information.
                  </li>
                  <li>
                    <strong>Delete Personal Data:</strong> You can request that
                    we delete specific personal data by contacting our support
                    team. Please note that this may affect your ability to use
                    certain features of our platform.
                  </li>
                  <li>
                    <strong>Object to Data Processing:</strong> You have the
                    right to object to our processing of your data. Please
                    contact our support team to discuss your concerns.
                  </li>
                  <li>
                    <strong>Request Data Portability:</strong> You have the right
                    to request that we transfer your data to another service.
                    Please contact our support team to discuss this option.
                  </li>
                  <li>
                    <strong>Withdraw Consent:</strong> You have the right to
                    withdraw your consent to our processing of your data at any
                    time. Please contact our support team to exercise this right.
                  </li>
                </ul>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">
                  7. Changes to This Privacy Policy
                </h2>
                <p>
                  We may update this policy periodically. Any changes will be
                  posted on this page, so please check back regularly.
                </p>
              </section>

              <section className="mb-8">
                <h2 className="text-2xl font-semibold mb-4">8. Contact Us</h2>
                <p>
                  If you have questions about this Privacy Policy, please
                  contact us at:
                </p>
                <p>
                  <strong>Email:</strong>{" "}
                  <a href="mailto:support@ametory.id" className="text-blue-600">
                    support@ametory.id
                  </a>
                </p>
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
export default PrivacyPolicy;
