package llmclient

const ExtractEmailAndNameSystemMessage = `
You will be given the transcript of a resume of a job applicant.
Please extract the email and name of the applicant from the resume. 
If there is no name or email of the resume's owner on the text fill it with either John doe or Jane Doe for the name and johndoe@gmail.com or janedoe@gmail.com for the email.`