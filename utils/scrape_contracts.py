from __future__ import print_function
import os.path
import hashlib
import base64
import pandas as pd
import numpy as np

# Google library imports
from googleapiclient.discovery import build
from google_auth_oauthlib.flow import InstalledAppFlow
from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials

# firestore library imports
import firebase_admin
from firebase_admin import credentials
from firebase_admin import firestore


# If modifying these scopes, delete the file token.json.
SCOPES = ['https://www.googleapis.com/auth/spreadsheets']

# The ID and range of a sample spreadsheet.
SAMPLE_SPREADSHEET_ID = '<ENTER SHEET ID>'
SAMPLE_RANGE_NAME = 'PlayerDatabase!C4:M514'

#League Name
LEAGUE = '<ENTER LEAGUE ID>'


def hashname(name) -> str:
    result = base64.urlsafe_b64encode(hashlib.md5(name.encode()).digest())
    string_res = result.decode()
    return string_res.replace("=", "")


def processContract(row, db):

    contractDetails = []

    for i in range(row[17]):
        if i == 0:
            x = {
                'Year': i + 1,
                'PaidAmount': row[16],
                'GuaranteedAmount': row[12],
                'TotalAmount': row[7]
            }
        else:
            x = {
                'Year': i + 1,
                'PaidAmount': 0,
                'GuaranteedAmount': row[12 + i],
                'TotalAmount': row[7 + i]
            }

        contractDetails.append(x)

    doc_ref = db.collection('leagues').document(LEAGUE).collection('playerContracts').document().set({
        "ContractDetails": contractDetails,
        'ContractLength': row[17],
        'CurrentYear': 2,
        'PlayerID': hashname(row[0]),
        'PlayerPosition': row[1],
        'RestructureStatus': 'ELIGIBLE',
        'TeamID': hashname(row[6]),
        'TotalContractValue': row[11],
        'TotalRemainingValue': row[20]
    })


def main():
    """Shows basic usage of the Sheets API.
    Prints values from a sample spreadsheet.
    """
    cred = credentials.Certificate('<cert.json HERE>')
    firebase_admin.initialize_app(cred)

    db = firestore.client()

    creds = None
    # The file token.json stores the user's access and refresh tokens, and is
    # created automatically when the authorization flow completes for the first
    # time.
    if os.path.exists('token.json'):
        creds = Credentials.from_authorized_user_file('token.json', SCOPES)
    # If there are no (valid) credentials available, let the user log in.
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(
                'credentials.json', SCOPES)
            creds = flow.run_local_server(port=0)
        # Save the credentials for the next run
        with open('token.json', 'w') as token:
            token.write(creds.to_json())

    service = build('sheets', 'v4', credentials=creds)

    # Call the Sheets API
    sheet = service.spreadsheets()
    result = sheet.values().get(spreadsheetId=SAMPLE_SPREADSHEET_ID,
                                range=SAMPLE_RANGE_NAME).execute()
    values = result.get('values', [])

    values_df = pd.DataFrame(np.row_stack(values))
    values_df.columns = ['playerNFL', 'playerPos', 'latestTrans', 'yearSigned', 'weekSigned',
                         'contractLength_str', 'owner', 'yr1value', 'yr2value', 'yr3value', 'yr4value']

    values_df['latestTrans'].fillna(0, inplace=True)

    values_df = values_df.astype({'yr1value': 'float', 'yr2value': 'float',
                                  'yr3value': 'float', 'yr4value': 'float'})

    values_df['yr1value'] = values_df.yr1value * 1000000
    values_df['yr2value'] = values_df.yr2value * 1000000
    values_df['yr3value'] = values_df.yr3value * 1000000
    values_df['yr4value'] = values_df.yr4value * 1000000

    values_df['totalContractValue'] = values_df.yr1value + values_df.yr2value + values_df.yr3value + values_df.yr4value
    values_df['yr1guaranteed'] = values_df.yr1value / 2
    values_df['yr2guaranteed'] = values_df.yr2value / 2
    values_df['yr3guaranteed'] = values_df.yr3value / 2
    values_df['yr4guaranteed'] = values_df.yr4value / 2
    values_df['yr1paidAmount'] = values_df.yr1value

#17
    values_df['contractLength'] = values_df.contractLength_str.apply(lambda x: int(x.split()[0]) if x != 'FA' else 'FA')
    values_df['restructureStatus'] = ['ELIGIBLE' if x != 'FA' else 'NA' for x in values_df.owner]
    values_df['contractYear'] = 2
    values_df['remainingValue'] = values_df.yr2value + values_df.yr3value + values_df.yr4value

    active_df = values_df[values_df.owner != 'FA']

    active_df.apply(lambda x: processContract(x, db), axis=1)

    print("contracts processed")


if __name__ == '__main__':
    main()