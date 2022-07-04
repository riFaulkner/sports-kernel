export const state = () => ({
    positionTypes: ["QB", "RB", "WR", "TE"],
    nflTeams: [
        {abbreviation: "ARI", longName: "Arizona Cardinals", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/ari.png&h=100&w=100"},
        {abbreviation: "ATL", longName: "Atlanta Falcons", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/atl.png&h=100&w=100"},
        {abbreviation: "BAL", longName: "Baltimore Ravens", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/bal.png&h=100&w=100"},
        {abbreviation: "BUF", longName: "Buffalo Bills", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/buf.png&h=100&w=100"},
        {abbreviation: "CAR", longName: "Carolina Panthers", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/car.png&h=100&w=100"},
        {abbreviation: "CHI", longName: "Chicago Bears",icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/chi.png&h=100&w=100"},
        {abbreviation: "CIN", longName: "Cincinnati Bengals", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/cin.png&h=100&w=100"},
        {abbreviation: "CLE", longName: "Cleavland Browns", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/cle.png&h=100&w=100"},
        {abbreviation: "DAL", longName: "Dallas Cowboys", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/dal.png&h=100&w=100"},
        {abbreviation: "DEN", longName: "Denver Broncos", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/den.png&h=100&w=100"},
        {abbreviation: "DET", longName: "Detroit Lions", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/det.png&h=100&w=100"},
        {abbreviation: "FA", longName: "Free Agent", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/leagues/500/nfl.png&w=100&h=100"},
        {abbreviation: "GB", longName: "Green Bay Packers", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/gb.png&h=100&w=100"},
        {abbreviation: "KC", longName: "Kansas City Chiefs", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/kc.png&h=100&w=100"},
        {abbreviation: "HOU", longName: "Houston Texans", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/hou.png&h=100&w=100"},
        {abbreviation: "IND", longName: "Indianapolis Colts", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/ind.png&h=100&w=100"},
        {abbreviation: "JAC", longName: "Jacksonville Jaguars", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/jax.png&h=100&w=100"},
        {abbreviation: "LAC", longName: "Los Angels Chargers", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/lac.png&h=100&w=100"},
        {abbreviation: "LAR", longName: "Los Angels Rams",icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/lar.png&h=100&w=100"},
        {abbreviation: "LV", longName: "Las Vegas Raiders", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/lv.png&h=100&w=100"},
        {abbreviation: "MIN", longName: "Minnesota Vikings", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/min.png&h=100&w=100"},
        {abbreviation: "MIA", longName: "Miami Dolphins", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/mia.png&h=100&w=100"},
        {abbreviation: "NE", longName: "New England Patriots", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/ne.png&h=100&w=100"},
        {abbreviation: "NO", longName: "New Orleans Saints",icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/no.png&h=100&w=100"},
        {abbreviation: "NYG", longName: "New York Giants", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/nyg.png&h=100&w=100"},
        {abbreviation: "NYJ", longName: "New York Jets", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/nyj.png&h=100&w=100"},
        {abbreviation: "PIT", longName: "Pittsburgh Steelers", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/pit.png&h=100&w=100"},
        {abbreviation: "PHI", longName: "Philadelphia Eagles", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/phi.png&h=100&w=100"},
        {abbreviation: "SF", longName: "San Francisco 49ers", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/sf.png&h=100&w=100"},
        {abbreviation: "SEA", longName: "Seattle Seahawks", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/sea.png&h=100&w=100"},
        {abbreviation: "TB", longName: "Tampa Bay Buccaneers", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/tb.png&h=100&w=100"},
        {abbreviation: "TEN", longName: "Tennessee Titans", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/ten.png&h=100&w=100"},
        {abbreviation: "WAS", longName: "Washington Commanders", icon: "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nfl/500/wsh.png&h=100&w=100"},
    ]
})

export const getters = {
    getPlayerPositionTypes(state) {
        return state.positionTypes
    },
    getNFLTeams(state) {
        return state.nflTeams
    }
}