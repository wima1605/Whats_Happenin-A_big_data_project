import React from 'react';
import './AboutUSPageStyles.css'; // Import your CSS file


function AboutUS() {
    return (
        <div>
            <nav className="navbar navbar-expand-lg bg-light">
                <div className="container-fluid">
                    <a className="navbar-brand me-auto" href="index.html">
                        <img src="../public/images/wh_logo.png" alt="" style={{ height: '100%', maxHeight: '50px' }} />
                    </a>
                    <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav">
                            <li className="nav-item">
                                <a className="nav-link" href="MVP.html">Home</a>
                            </li>
                            <li className="nav-item">
                                <a className="nav-link active" aria-current="page" href="About_US.html">About Us</a>
                            </li>
                            <li className="nav-item">
                                <a className="nav-link" href="Contact_US.html">Contact Us</a>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
            <div className="container">
                <div className="row">
                    <div className="col-md-4 mb-4">
                        <div className="profile">
                            <a href="https://www.linkedin.com/in/shanya-chaubey-2a663169/" target="_blank" rel="noreferrer" title="Visit Shanya Chaubey's LinkedIn profile">
                                <img src="https://media.licdn.com/dms/image/C4E03AQHGL1-GHi4wPg/profile-displayphoto-shrink_400_400/0/1569849280789?e=1716422400&v=beta&t=br-QISY-o11VLytT_Voz_wu-VcZWa-C6LPtskFKy778" alt="shanya" width="400" height="auto" />
                            </a>
                            <h3>Shanya Chaubey</h3>
                            <p>Graduate Student @ Colorado Boulder. Project Manager. Head of Machine Learning.</p>
                        </div>
                    </div>
                    <div className="col-md-4 mb-4">
                        <div className="profile">
                            <a href="https://www.linkedin.com/in/bill-martynowicz-12895121a/" target="_blank" rel="noreferrer" title="Visit Bill Martynowicz's LinkedIn profile">
                                <img src="https://media.licdn.com/dms/image/D5603AQHXT41B-Eqg0Q/profile-displayphoto-shrink_800_800/0/1707510291653?e=1716422400&v=beta&t=i_QctShT-drj5u3vUUIhDbJlO1SXoKVIWssaQ9-OkFE" alt="bill" width="400" height="auto" />
                            </a>
                            <h3>Bill Martynowicz</h3>
                            <p>Junior CS @ Colorado Boulder. Full-Stack Software Engineer.</p>
                        </div>
                    </div>
                    <div className="col-md-4 mb-4">
                        <div className="profile">
                            <img src="https://assets.nick.com/uri/mgid:arc:imageassetref:shared.nick.us:a625d441-bbbf-42c8-9927-6a0157aac911?quality=0.7&gen=ntrn&legacyStatusCode=true" alt="prateek" width="400" height="auto" />
                            <h3>Prateek Kumar</h3>
                            <p>Graduate Student @ Colorado Boulder. Full-Stack Software Engineer. Cloud Engineer. Data Engineer.</p>
                        </div>
                    </div>
                    <div className="col-md-4 mb-4">
                        <div className="profile">
                            <a href="https://www.linkedin.com/in/justinjamrowski" target="_blank" rel="noreferrer" title="Visit Justin Jamrowski's LinkedIn profile">
                                <img src="https://media.licdn.com/dms/image/D5603AQFwwH1EviPS5g/profile-displayphoto-shrink_800_800/0/1707370144433?e=1716422400&v=beta&t=vM5t-1w0QypstfXTQ2JDv3efdoQVsJ3FYHNE-4ieoxM" alt="justin" width="400" height="auto" />
                            </a>
                            <h3>Justin Jamrowski</h3>
                            <p>Junior CS @ Colorado Boulder. Co-Project Manager. Full-Stack Software Engineer. UX Designer.</p>
                        </div>
                    </div>
                    <div className="col-md-4 mb-4">
                        <div className="profile">
                            <a href="https://www.linkedin.com/in/niharika-n-govinda-0689461aa/" target="_blank" rel="noreferrer" title="Visit Niharika N Govinda's LinkedIn profile">
                                <img src="https://media.licdn.com/dms/image/D5635AQEd0GP3z3hJVw/profile-framedphoto-shrink_400_400/0/1702072467984?e=1711411200&v=beta&t=1eLMcMCfgK6-7xlHqv-Nu9-2DMJyoVb3aYf6Vx-gg7E" alt="niharika" width="400" height="auto" />
                            </a>
                            <h3>Niharika N Govinda</h3>
                            <p>Graduate Student @ Colorado Boulder. Machine Learning Engineer.</p>
                        </div>
                    </div>
                    <div className="col-md-4 mb-4">
                        <div className="profile">
                            <a href="https://www.linkedin.com/in/lohith-ramesh-9094941b1/" target="_blank" rel="noreferrer" title="Visit Lohith Ramesh's LinkedIn profile">
                                <img src="https://media.licdn.com/dms/image/D5635AQHuO6Xdc0-pTw/profile-framedphoto-shrink_400_400/0/1706418740703?e=1711407600&v=beta&t=rOtmB2O4YFCU6UGXg7wKLbm8n2zSrGCAeOntoKB-ows" alt="lohith" width="400" height="auto" />
                            </a>
                            <h3>Lohith Ramesh</h3>
                            <p>Graduate Student @ Colorado Boulder. Machine Learning Engineer.</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default AboutUS;