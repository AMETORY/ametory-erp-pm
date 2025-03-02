import * as CryptoJS from 'crypto-js';

export const asyncStorage = {
  async setItem(key, value) {
    return new Promise((resolve, reject) => {
      try {
        localStorage.setItem(key, encrypt(value));
        resolve();
      } catch (error) {
        reject(error);
      }
    });
  },

  async getItem(key) {
    return new Promise((resolve, reject) => {
      try {
        const value = localStorage.getItem(key);
        if (!value) {
          resolve(null);
          return;
        }
        resolve(decrypt(value));
      } catch (error) {
        reject(error);
      }
    });
  },

  async removeItem(key) {
    return new Promise((resolve, reject) => {
      try {
        localStorage.removeItem(key);
        resolve();
      } catch (error) {
        reject(error);
      }
    });
  },
};

function encrypt(txt) {
  return CryptoJS.AES.encrypt(txt, process.env.REACT_APP_SECRET_KEY).toString();
}

function decrypt(txtToDecrypt) {
  if (!txtToDecrypt) return
  return CryptoJS.AES.decrypt(txtToDecrypt, process.env.REACT_APP_SECRET_KEY).toString(CryptoJS.enc.Utf8);
}
