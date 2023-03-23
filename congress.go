package main

type LatestAction struct {
    ActionDate string `json:"actionDate"`
    Text string `json:"text"`
}

type Bill struct {
    Actions struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"actions"`
    Amendments struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"amendments"`
    CBOCostEstimates []struct {
        Description string `json:"description"`
        PubDate     string `json:"pubDate"`
        Title       string `json:"title"`
        URL         string `json:"url"`
    } `json:"cboCostEstimates"`
    CommitteeReports []struct {
        Citation string `json:"citation"`
        URL      string `json:"url"`
    } `json:"committeeReports"`
    Committees struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"committees"`
    Congress                        int    `json:"congress"`
    ConstitutionalAuthorityStatementText string `json:"constitutionalAuthorityStatementText"`
    Cosponsors struct {
        Count                           int    `json:"count"`
        CountIncludingWithdrawnCosponsors int    `json:"countIncludingWithdrawnCosponsors"`
        URL                             string `json:"url"`
    } `json:"cosponsors"`
    IntroducedDate                  string `json:"introducedDate"`
    LatestAction struct {
        ActionDate string `json:"actionDate"`
        Text       string `json:"text"`
    } `json:"latestAction"`
    Laws []struct {
        Number string `json:"number"`
        Type   string `json:"type"`
    } `json:"laws"`
    Number       string `json:"number"`
    OriginChamber string `json:"originChamber"`
    PolicyArea   struct {
        Name string `json:"name"`
    } `json:"policyArea"`
    RelatedBills struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"relatedBills"`
    Sponsors []struct {
        BioguideId  string `json:"bioguideId"`
        District   int    `json:"district"`
        FirstName  string `json:"firstName"`
        FullName   string `json:"fullName"`
        IsByRequest string `json:"isByRequest"`
        LastName   string `json:"lastName"`
        MiddleName string `json:"middleName"`
        Party      string `json:"party"`
        State      string `json:"state"`
        URL        string `json:"url"`
    } `json:"sponsors"`
    Subjects struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"subjects"`
    Summaries struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"summaries"`
    TextVersions struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"textVersions"`
    Title  string `json:"title"`
    Titles struct {
        Count int    `json:"count"`
        URL   string `json:"url"`
    } `json:"titles"`
    Type   string `json:"type"`
    UpdateDate              string `json:"updateDate"`
    UpdateDateIncludingText string `json:"updateDateIncludingText"`
}

type Pagination struct {
    Count int `json:"count"`
    Next string `json:"next"`
}

type Request struct {
    ContentType string `json:"contentType"`
    Format string `json:"format"`
}

type CongressBills struct {
    Bills []Bill `json:"bills"`
    Pagination Pagination `json:"pagination"`
    Request Request `json:"request"`
}

type BillRequest struct {
    Bill Bill `json:"bill"`
}

type TextRequest struct {
    Pagination struct {
        Count int `json:"count"`
    } `json:"pagination"`
    Request struct {
        BillNumber string `json:"billNumber"`
        BillType   string `json:"billType"`
        BillURL    string `json:"billUrl"`
        Congress   string `json:"congress"`
        ContentType string `json:"contentType"`
        Format     string `json:"format"`
    } `json:"request"`
    TextVersions []struct {
        Date    interface{} `json:"date"`
        Formats []string    `json:"formats"`
        Type    interface{} `json:"type"`
    } `json:"textVersions"`
}
